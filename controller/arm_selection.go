package controller

import (
	"github.com/gin-gonic/gin"
	"jackpot-mab/arm-selector/id"
	"jackpot-mab/arm-selector/policy"
	"net/http"
)

// @BasePath /api/v1

// ArmSelectionController godoc
// @Summary select arm
// @Schemes
// @Description Select the arm based on the experiment policy.
// @Tags arm-selector
// @Accept json
// @Param context body policy.Context true "Context Data"
// @Produce json
// @Success 200 {string} Arm
// @Router /arm/selection [post]
func ArmSelectionController(g *gin.Context) {

	var decisionContext policy.Context
	if err := g.BindJSON(&decisionContext); err != nil {
		return
	}

	// Get experiment data from external service
	mockData := policy.Experiment{
		Id:     "1-EEE-3",
		Type:   "epsilon_greedy",
		Arms:   []policy.Arm{{Name: "A"}, {Name: "B"}, {Name: "C"}},
		Params: map[string][]interface{}{"alpha": {0.2}},
	}

	// Get the value returned by the model
	mockExpectedReward := []policy.ExpectedReward{
		{
			Arm:   policy.Arm{Name: "A"},
			Pulls: 1,
			Value: 102,
		},
		{
			Arm:   policy.Arm{Name: "B"},
			Pulls: 3,
			Value: 305,
		},
		{
			Arm:   policy.Arm{Name: "C"},
			Pulls: 12,
			Value: 308,
		}}

	currentPolicy, _ := policy.MakePolicy(mockData)

	armSelected := currentPolicy.SelectArm(mockExpectedReward)

	randomID, _ := id.GenerateRandomID()

	r := Response{
		DecisionId: randomID,
		Arm:        *armSelected,
	}

	g.JSON(http.StatusOK, r)

}
