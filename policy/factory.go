package policy

import "errors"

const (
	EpsilonGreedyType = "epsilon_greedy"
	Ucb1Type          = "ucb1"
)

func MakePolicy(experimentData Experiment) (Policy, error) {

	if experimentData.PolicyType == Ucb1Type {
		return MakeUCB1(), nil
	}

	if experimentData.PolicyType == EpsilonGreedyType {

		return MakeEpsilonGreedy(experimentData.Parameters)

	}

	return nil, errors.New("experiment type not recognized")
}
