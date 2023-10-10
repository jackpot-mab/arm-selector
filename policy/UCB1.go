package policy

import (
	"math"
	"strconv"
)

// UCB1 is a policy that computes a mean_reward or predicted_reward for a context,
// along with a confidence radius. Then the algorithm selects the higher mean_reward + confidence_radius
// arm.
type UCB1 struct {
	// Se puede agregar un parámetro de regularización configurable.
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
	radius := math.Sqrt(2 * math.Log(float64(totalPulls)) / getPullsParams(arm.Arm.RewardDataParameters))
	return arm.Value + radius
}

func getTotalPulls(armsExpectedRewards []ExpectedReward) int {
	sum := 0
	for _, a := range armsExpectedRewards {
		pulls := getPullsParams(a.Arm.RewardDataParameters)
		sum = int(float64(sum) + pulls)
	}
	return sum
}

func getPullsParams(parameters []RewardDataParameter) float64 {
	pullsStr := selectParameter("pulls", parameters)
	if pullsStr == "" {
		return 0
	}
	pulls, _ := strconv.ParseFloat(pullsStr, 64)
	return pulls
}
