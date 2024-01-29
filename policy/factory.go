package policy

import "errors"

const (
	EpsilonGreedyType    = "epsilon_greedy"
	Ucb1Type             = "ucb1"
	ThompsonSamplingType = "thompson_sampling"
)

func MakePolicy(experimentData Experiment) (Policy, error) {

	if experimentData.PolicyType == Ucb1Type {
		return MakeUCB1(), nil
	}

	if experimentData.PolicyType == EpsilonGreedyType {

		return MakeEpsilonGreedy(experimentData.Parameters)

	}

	if experimentData.PolicyType == ThompsonSamplingType {

		return MakeThompsonSampling([]float32{0, 1}), nil

	}

	return nil, errors.New("experiment type not recognized")
}
