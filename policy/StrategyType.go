package policy

// StrategyType enum of policy strategies available
type StrategyType uint8

const (
	EpsilonGreedyType StrategyType = iota
	UCB1Type
)

func (d StrategyType) String() string {
	switch d {
	case UCB1Type:
		return "ucb1"
	case EpsilonGreedyType:
		return "epsilon_greedy"
	default:
		return "invalid_strategy_type"
	}
}
