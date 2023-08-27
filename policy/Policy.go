package policy

type Experiment struct {
	Id     string
	Type   string
	Arms   []Arm
	Params map[string][]interface{}
}

type Arm struct {
	Name string
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
