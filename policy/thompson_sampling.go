package policy

import (
	"errors"
	"math/rand"
	"time"
)

// ThompsonSampling is a policy that computes the expected reward based on the probability
// distribution returned by the predictor.
type ThompsonSampling struct {
	rewardValues []float32
	seed         *int64
}

func MakeThompsonSampling(rewardValues []float32) *ThompsonSampling {
	return &ThompsonSampling{rewardValues: rewardValues}
}

func (t *ThompsonSampling) SelectArm(arms []Arm, armsExpectedRewards []ExpectedReward) *Arm {

	if len(arms) == 0 {
		return nil
	}

	if len(armsExpectedRewards) == 0 {
		randomIndex := rand.Intn(len(arms))
		return &arms[randomIndex]
	}

	seed := time.Now().UnixNano()
	if t.seed != nil {
		seed = *t.seed
	}

	maxReward := float32(-1)
	var armWithMaxReward *Arm

	for _, armExpectedRewards := range armsExpectedRewards {
		arm := armExpectedRewards.Arm
		reward, err := t.sampleRewardFromProbaDistribution(t.rewardValues, armExpectedRewards.Probabilities)

		if err != nil {
			continue
		}

		if reward == maxReward {
			withSeed := rand.New(rand.NewSource(seed))
			r := withSeed.Float32()
			if r >= 0.5 {
				maxReward = reward
				armWithMaxReward = &arm
			}
		}

		if reward > maxReward {
			maxReward = reward
			armWithMaxReward = &arm
		}
	}

	return armWithMaxReward
}

func (t *ThompsonSampling) sampleRewardFromProbaDistribution(values []float32, probas []float32) (float32, error) {

	seed := time.Now().UnixNano()
	if t.seed != nil {
		seed = *t.seed
	}

	withSeed := rand.New(rand.NewSource(seed))
	r := withSeed.Float32()

	cumulativeProb := float32(0)
	for i, prob := range probas {
		cumulativeProb += prob
		if r <= cumulativeProb {
			return values[i], nil
		}
	}

	// This should not be reached, but if it does, return 0 reward
	return 0, errors.New("not matching probas with possible values")
}
