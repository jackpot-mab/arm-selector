package policy

import (
	"errors"
	"math/rand"
	"time"
)

type EpsilonGreedy struct {
	alpha float64
	seed  *int64
}

func MakeEpsilonGreedy(experimentParams map[string]interface{}) (*EpsilonGreedy, error) {

	alpha, ok := experimentParams["alpha"].(float64)
	if !ok {
		return &EpsilonGreedy{alpha: 0.5}, errors.New("missing parameter alpha")
	}

	if experimentParams["seed"] != nil {
		seed := experimentParams["seed"].(int64)
		return &EpsilonGreedy{alpha: alpha, seed: &seed}, nil
	}

	return &EpsilonGreedy{alpha: alpha}, nil
}

func (e *EpsilonGreedy) SelectArm(arms []Arm, armsExpectedRewards []ExpectedReward) *Arm {

	seed := time.Now().UnixNano()
	if e.seed != nil {
		seed = *e.seed
	}

	if len(armsExpectedRewards) == 0 && len(arms) > 0 {
		randomIndex := rand.Intn(len(arms))
		return &arms[randomIndex]
	}

	// rolls a dice to explore or to exploit
	r := rand.New(rand.NewSource(seed))
	dice := r.Float64()

	var selection *Arm

	if dice < e.alpha {
		// explore
		selection = getRandomArm(armsExpectedRewards, r)
	} else {
		// EXPLOIT
		selection = getMaxRewardArm(armsExpectedRewards)
	}

	return selection
}

func getMaxRewardArm(armsExpectedRewards []ExpectedReward) *Arm {

	if len(armsExpectedRewards) == 0 {
		return nil
	}

	selectedArm := armsExpectedRewards[0]
	for _, armReward := range armsExpectedRewards[1:] {
		if armReward.Value > selectedArm.Value {
			selectedArm = armReward
		}
	}

	return &selectedArm.Arm
}

func getRandomArm(armsExpectedRewards []ExpectedReward, r *rand.Rand) *Arm {
	if len(armsExpectedRewards) == 0 {
		return nil
	}

	randomIndex := r.Intn(len(armsExpectedRewards))
	return &armsExpectedRewards[randomIndex].Arm
}
