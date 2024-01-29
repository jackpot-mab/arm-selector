package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"jackpot-mab/arm-selector/policy"
	"net/http"
	"time"
)

type RewardPredictorService interface {
	GetRewardPrediction(experimentId string,
		arm policy.Arm, features []interface{},
		classes []string) (policy.ExpectedReward, error)
}

type RewardPredictorServiceImpl struct {
	client *http.Client
	url    string
}

func MakeRewardPredictorService(url string, timeoutMillis int) RewardPredictorService {
	return &RewardPredictorServiceImpl{
		client: &http.Client{
			Timeout: time.Duration(timeoutMillis) * time.Millisecond,
		},
		url: url,
	}
}

type RewardPredictorResponse struct {
	Prediction    float64   `json:"prediction"`
	Probabilities []float32 `json:"probabilities"`
}

type RewardPredictorRequest struct {
	Context []interface{} `json:"context"`
	Model   string        `json:"model"`
	Sample  bool          `json:"sample"`
	Classes []string      `json:"classes"`
}

func (rp *RewardPredictorServiceImpl) GetRewardPrediction(
	experimentId string, arm policy.Arm, features []interface{},
	outputClasses []string) (policy.ExpectedReward, error) {

	body := RewardPredictorRequest{
		Context: features,
		Model:   fmt.Sprintf("%s:%s", experimentId, arm.Name),
		Classes: outputClasses,
	}

	bodyMarshaled, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", rp.url, bytes.NewBuffer(bodyMarshaled))
	if err != nil {
		return policy.ExpectedReward{}, err
	}

	resp, err := rp.client.Do(req)
	if err != nil {
		return policy.ExpectedReward{}, err
	}
	defer resp.Body.Close()

	jsonResponse, _ := io.ReadAll(resp.Body)

	var output RewardPredictorResponse
	err = json.Unmarshal(jsonResponse, &output)
	if err != nil {
		return policy.ExpectedReward{}, err
	}

	return TransformToExpectedReward(output, arm), nil
}

func TransformToExpectedReward(response RewardPredictorResponse, arm policy.Arm) policy.ExpectedReward {
	return policy.ExpectedReward{
		Arm:           arm,
		Value:         response.Prediction,
		Probabilities: response.Probabilities,
	}
}
