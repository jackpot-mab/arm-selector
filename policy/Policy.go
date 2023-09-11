package policy

type Experiment struct {
	ExperimentId string                 `json:"experiment_id"`
	PolicyType   string                 `json:"policy_type"`
	Arms         []Arm                  `json:"arms"`
	Parameters   map[string]interface{} `json:"parameters"`
}

type Arm struct {
	Name string `json:"name"`
}

type ExpectedReward struct {
	Pulls int
	Arm   Arm
	Value float64 // TODO evaluar bien qué necesitamos del modelo.
}

type Context map[string]interface{}

type Policy interface {
	SelectArm(armsExpectedRewards []ExpectedReward) *Arm
}
