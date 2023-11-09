package controller

import (
	"github.com/gin-gonic/gin"
	"jackpot-mab/arm-selector/id"
	"jackpot-mab/arm-selector/policy"
	"jackpot-mab/arm-selector/service"
	"net/http"
	"strings"
)

type ArmSelectorController struct {
	ExperimentsParamsService service.ExperimentParamsService
	RewardPredictorService   service.RewardPredictorService
}

// @BasePath /api/v1

// ArmSelectionController godoc
// @Summary select arm
// @Schemes
// @Description Select the arm based on the experiment policy.
// @Tags arm-selector
// @Accept json
// @Param experiment_id path string true "ID of the current experiment."
// @Param context body policy.Context true "Context Data"
// @Produce json
// @Success 200 {string} Arm
// @Router /arm/selection/{experiment_id} [post]
func (a *ArmSelectorController) ArmSelectionController(g *gin.Context) {

	var decisionContext policy.Context
	if err := g.BindJSON(&decisionContext); err != nil {
		g.JSON(http.StatusInternalServerError, err)
		return
	}

	experimentId := g.Param("experiment_id")

	// Get experiment data from external service
	experimentData, err := a.ExperimentsParamsService.GetExperiment(experimentId)

	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error obtaining experiment from " +
			"experiment-params service, check if the experiment exists and its config is correct. Error: " + err.Error()})
		return
	}

	modelPredictions := service.GetMultipleRewardPredictionsParallel(
		experimentId,
		experimentData.Arms,
		deduceInputFeatures(decisionContext, experimentData.ModelParams.InputFeatures),
		experimentData.ModelParams.OutputClasses,
		a.RewardPredictorService)

	currentPolicy, _ := policy.MakePolicy(experimentData)

	armSelected := currentPolicy.SelectArm(experimentData.Arms, modelPredictions)

	randomID, _ := id.GenerateRandomID()

	r := Response{
		DecisionId: randomID,
		Arm:        *armSelected,
	}

	g.JSON(http.StatusOK, r)

}

func deduceInputFeatures(context policy.Context, inputFeatures []string) []interface{} {
	var inputFeaturesContext []interface{}
	for _, feature := range inputFeatures {
		// TODO Check input features exist
		inputFeaturesContext = append(inputFeaturesContext, context[strings.ToLower(feature)])
	}
	return inputFeaturesContext
}
