package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"jackpot-mab/arm-selector/controller"
	"jackpot-mab/arm-selector/docs"
	"jackpot-mab/arm-selector/service"
	"log"
	"net/http"
	"os"
	"strconv"
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func healthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "jackpot-mab:arm-selector")
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	armSelectionCtrl := controller.ArmSelectorController{
		ExperimentsParamsService: InitExperimentService(),
		RewardPredictorService:   InitRewardPredictorService()}

	docs.SwaggerInfo.BasePath = "/api/v1"
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/arm")
		{
			eg.POST("/selection/:experiment_id", armSelectionCtrl.ArmSelectionController)
		}
	}

	router.GET("/metrics", prometheusHandler())
	router.GET("/", healthCheck)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("0.0.0.0:8090")
}

func InitExperimentService() service.ExperimentParamsService {
	ExperimentsParamsRepoUrl := os.Getenv("EXPERIMENTS_PARAMS_SERVICE_URL")
	TimeoutMillis, err := strconv.Atoi(
		os.Getenv("EXPERIMENTS_PARAMS_SERVICE_TIMEOUT_MILLIS"))

	if err != nil {
		log.Fatalf("error reading env variable EXPERIMENTS_PARAMS_SERVICE_TIMEOUT_MILLIS")
	}

	return service.MakeExperimentsParamsService(ExperimentsParamsRepoUrl, TimeoutMillis)
}

func InitRewardPredictorService() service.RewardPredictorService {
	RewardPredictorUrl := os.Getenv("REWARD_PREDICTOR_URL")
	TimeoutMillis, err := strconv.Atoi(
		os.Getenv("REWARD_PREDICTOR_TIMEOUT_MILLIS"))

	if err != nil {
		log.Fatalf("error reading env variable REWARD_PREDICTOR_TIMEOUT_MILLIS")
	}

	return service.MakeRewardPredictorService(RewardPredictorUrl, TimeoutMillis)
}
