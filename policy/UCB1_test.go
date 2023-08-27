package policy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUCB1_SelectArm(t *testing.T) {

	var tests = []struct {
		name        string
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
			armsRewards: []ExpectedReward{
				{
					Arm:   Arm{Name: "A"},
					Pulls: 1,
					Value: 102,
				},
				{
					Arm:   Arm{Name: "B"},
					Pulls: 3,
					Value: 305,
				},
				{
					Arm:   Arm{Name: "C"},
					Pulls: 12,
					Value: 308,
				},
			},
			selection: &Arm{Name: "C"},
		},
	}

	for _, tt := range tests {
		t.Logf("Executing test => %v", tt.name)

		policy := UCB1{}

		armSelected := policy.SelectArm(tt.armsRewards)

		assert.Equal(t, tt.selection, armSelected)

	}
}
