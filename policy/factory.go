package policy

import "errors"

const (
	EpsilonGreedyType = "epsilon_greedy"
	Ucb1Type          = "ucb1"
)

func MakePolicy(experimentData Experiment) (Policy, error) {

	if experimentData.Type == Ucb1Type {
		return MakeUCB1(), nil
	}

	if experimentData.Type == EpsilonGreedyType {

		return MakeEpsilonGreedy(experimentData.Params)

	}

	return nil, errors.New("experiment type not recognized")
}
