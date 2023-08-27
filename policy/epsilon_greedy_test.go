package policy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEpsilonGreedy_SelectArm(t *testing.T) {

	var tests = []struct {
		name        string
		seed        int64
		alpha       float32
		armsRewards []ExpectedReward
		selection   *Arm
	}{
		{
			name:        "No arms",
			seed:        42,
			alpha:       0.2,
			armsRewards: []ExpectedReward{},
			selection:   nil,
		},
		{
			name:  "Select arm with highest reward",
			seed:  42,
			alpha: 0.2,
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
		{
			name:  "Select a random arm",
			seed:  131,
			alpha: 0.2,
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
			selection: &Arm{Name: "A"},
		},
	}

	for _, tt := range tests {
		t.Logf("Executing test => %v", tt.name)

		policy := EpsilonGreedy{
			alpha: tt.alpha,
			seed:  &tt.seed,
		}

		armSelected := policy.SelectArm(tt.armsRewards)

		assert.Equal(t, tt.selection, armSelected)

	}
}
