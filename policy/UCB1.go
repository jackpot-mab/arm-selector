package policy

import "math"

type UCB1 struct {
	// PARAMS?
}

func MakeUCB1() Policy {
	return &UCB1{}
}

func (e *UCB1) SelectArm(armsExpectedRewards []ExpectedReward) *Arm {

	if len(armsExpectedRewards) == 0 {
		return nil
	}

	totalPulls := getTotalPulls(armsExpectedRewards)

	if totalPulls == 0 {
		return &armsExpectedRewards[0].Arm
	}

	maxUCB := getArmUCB(armsExpectedRewards[0], totalPulls)
	selectedArm := armsExpectedRewards[0]

	// For each arm calc radius plus mean
	for _, arm := range armsExpectedRewards[1:] {
		armUcb := getArmUCB(arm, totalPulls)
		if armUcb >= maxUCB {
			selectedArm = arm
			maxUCB = armUcb
		}
	}

	return &selectedArm.Arm
}

func getArmUCB(arm ExpectedReward, totalPulls int) float64 {
	radius := math.Sqrt(2 * math.Log(float64(totalPulls)) / float64(arm.Pulls))
	return arm.Value + radius
}

func getTotalPulls(armsExpectedRewards []ExpectedReward) int {
	sum := 0
	for _, a := range armsExpectedRewards {
		sum = sum + a.Pulls
	}
	return sum
}
