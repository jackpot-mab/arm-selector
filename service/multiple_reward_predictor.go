package service

import (
	"jackpot-mab/arm-selector/policy"
	"log"
	"sync"
)

func GetMultipleRewardPredictionsParallel(experimentId string,
	arms []policy.Arm, context []interface{}, classes []string, service RewardPredictorService) []policy.ExpectedReward {

	parallelRequests := len(arms)
	var results []policy.ExpectedReward

	var wg sync.WaitGroup
	resultChan := make(chan struct {
		policy.ExpectedReward
		error
	}, parallelRequests)

	for _, arm := range arms {
		wg.Add(1)
		go func(arm policy.Arm) {
			defer wg.Done()
			result, err := service.GetRewardPrediction(experimentId, arm, context, classes)
			resultChan <- struct {
				policy.ExpectedReward
				error
			}{result, err}
		}(arm)
	}

	wg.Wait()
	close(resultChan)

	for result := range resultChan {
		if result.error != nil {
			log.Printf("Error getting reward prediction for arms.")
			print(result.error.Error())
			return []policy.ExpectedReward{}
		} else {
			results = append(results, result.ExpectedReward)
		}
	}

	return results

}
