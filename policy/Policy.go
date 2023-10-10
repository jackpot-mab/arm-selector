package policy

import "fmt"

type Experiment struct {
	ExperimentId string                 `json:"experiment_id"`
	PolicyType   string                 `json:"policy_type"`
	Arms         []Arm                  `json:"arms"`
	Parameters   map[string]interface{} `json:"parameters"`
	ModelParams  MLModelParameters      `json:"model_parameters"`
}

type MLModelParameters struct {
	ModelType     string   `json:"model_type"`
	InputFeatures []string `json:"input_features"`
	OutputClasses []string `json:"output_classes"`
}

type Arm struct {
	Name                 string                `json:"name"`
	RewardDataParameters []RewardDataParameter `json:"reward_data_parameters"`
}

type RewardDataParameter struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type ExpectedReward struct {
	Arm   Arm
	Value float64
}

type Context map[string]interface{}

type Policy interface {
	SelectArm(armsExpectedRewards []ExpectedReward) *Arm
}

func selectParameter(paramName string, rewardDataParameter []RewardDataParameter) string {
	for _, p := range rewardDataParameter {
		if paramName == p.Name {
			return fmt.Sprintf("%v", p.Value)
		}
	}
	return ""
}
