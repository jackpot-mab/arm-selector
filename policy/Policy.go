package policy

type Experiment struct {
	Id   string
	Type StrategyType
	Arms []Arm
}

type Arm struct {
	Name string
}

type Context map[string]string

type Policy interface {
	MakeDecision(exp Experiment, decisionContext Context) Arm
}
