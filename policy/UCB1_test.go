package policy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUCB1_SelectArm(t *testing.T) {

	var tests = []struct {
		name        string
		arms        []Arm
		armsRewards []ExpectedReward
		selection   *Arm
	}{
		{
			name:        "No arms",
			armsRewards: []ExpectedReward{},
			selection:   nil,
		},
		{
			name: "Select arm with highest reward",
			arms: []Arm{{
				Name:                 "A",
				RewardDataParameters: []RewardDataParameter{{Name: "pulls", Value: 1}}},
				{
					Name:                 "B",
					RewardDataParameters: []RewardDataParameter{{Name: "pulls", Value: 3}},
				},
				{
					Name:                 "C",
					RewardDataParameters: []RewardDataParameter{{Name: "pulls", Value: 12}},
				},
			},
			armsRewards: []ExpectedReward{
				{
					Arm:   Arm{Name: "A", RewardDataParameters: []RewardDataParameter{{Name: "pulls", Value: 1}}},
					Value: 102,
				},
				{
					Arm:   Arm{Name: "B", RewardDataParameters: []RewardDataParameter{{Name: "pulls", Value: 3}}},
					Value: 305,
				},
				{
					Arm:   Arm{Name: "C", RewardDataParameters: []RewardDataParameter{{Name: "pulls", Value: 12}}},
					Value: 308,
				},
			},
			selection: &Arm{Name: "C", RewardDataParameters: []RewardDataParameter{{Name: "pulls", Value: 12}}},
		},
	}

	for _, tt := range tests {
		t.Logf("Executing test => %v", tt.name)

		policy := UCB1{}

		armSelected := policy.SelectArm(tt.arms, tt.armsRewards)

		assert.Equal(t, tt.selection, armSelected)

	}
}
