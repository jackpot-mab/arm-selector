package policy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEpsilonGreedy_SelectArm(t *testing.T) {

	var tests = []struct {
		name        string
		arms        []Arm
		seed        int64
		alpha       float32
		armsRewards []ExpectedReward
		selection   *Arm
	}{
		{
			name:        "No arms",
			arms:        []Arm{},
			seed:        42,
			alpha:       0.2,
			armsRewards: []ExpectedReward{},
			selection:   nil,
		},
		{
			name:  "Select arm with highest reward",
			arms:  []Arm{{Name: "A"}, {Name: "B"}, {Name: "C"}},
			seed:  42,
			alpha: 0.2,
			armsRewards: []ExpectedReward{
				{
					Arm:   Arm{Name: "A"},
					Value: 102,
				},
				{
					Arm:   Arm{Name: "B"},
					Value: 305,
				},
				{
					Arm:   Arm{Name: "C"},
					Value: 308,
				},
			},
			selection: &Arm{Name: "C"},
		},
		{
			name:  "Select a random arm",
			arms:  []Arm{{Name: "A"}, {Name: "B"}, {Name: "C"}},
			seed:  131,
			alpha: 0.2,
			armsRewards: []ExpectedReward{
				{
					Arm:   Arm{Name: "A"},
					Value: 102,
				},
				{
					Arm:   Arm{Name: "B"},
					Value: 305,
				},
				{
					Arm:   Arm{Name: "C"},
					Value: 308,
				},
			},
			selection: &Arm{Name: "A"},
		},
	}

	for _, tt := range tests {
		t.Logf("Executing test => %v", tt.name)

		policy := EpsilonGreedy{
			alpha: float64(tt.alpha),
			seed:  &tt.seed,
		}

		armSelected := policy.SelectArm(tt.arms, tt.armsRewards)

		assert.Equal(t, tt.selection, armSelected)

	}
}
