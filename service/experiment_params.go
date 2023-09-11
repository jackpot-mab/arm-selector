package service

import (
	"encoding/json"
	"errors"
	"io"
	"jackpot-mab/arm-selector/policy"
	"net/http"
	"time"
)

type ExperimentParamsService interface {
	GetExperiment(experimentId string) (policy.Experiment, error)
}

type ExperimentParamsServiceImpl struct {
	client *http.Client
	url    string
}

func MakeExperimentsParamsService(url string, timeoutMillis int) ExperimentParamsService {
	return &ExperimentParamsServiceImpl{
		client: &http.Client{
			Timeout: time.Duration(timeoutMillis) * time.Millisecond,
		},
		url: url,
	}
}

func (e *ExperimentParamsServiceImpl) GetExperiment(experimentId string) (policy.Experiment, error) {

	req, err := http.NewRequest("GET", e.url+experimentId, nil)
	if err != nil {
		return policy.Experiment{}, err
	}

	resp, err := e.client.Do(req)
	if err != nil {
		return policy.Experiment{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return policy.Experiment{}, err
	}

	var experiment policy.Experiment
	err = json.Unmarshal(body, &experiment)
	if err != nil {
		return policy.Experiment{}, err
	}

	if experiment.ExperimentId == "" {
		return policy.Experiment{}, errors.New("experiment not found")
	}

	return experiment, nil

}
