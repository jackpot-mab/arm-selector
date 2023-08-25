package policy

type EpsilonGreedy struct {
	alfa float32
}

func (e *EpsilonGreedy) MakeDecision(exp Experiment, decisionContext Context) Arm {
	return Arm{}
}
