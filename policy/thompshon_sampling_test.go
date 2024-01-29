package policy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThompsonSampling_SelectArm(t *testing.T) {

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
				Name: "A",
			},
				{
					Name: "B",
				},
				{
					Name: "C",
				},
			},
			armsRewards: []ExpectedReward{
				{
					Arm:           Arm{Name: "A"},
					Value:         1,
					Probabilities: []float32{0.9, 1},
				},
				{
					Arm:           Arm{Name: "B"},
					Value:         0,
					Probabilities: []float32{0.2, 0.8},
				},
				{
					Arm:           Arm{Name: "C"},
					Value:         1,
					Probabilities: []float32{0.5, 0.5},
				},
			},
			selection: &Arm{Name: "A"},
		},
	}

	for _, tt := range tests {
		t.Logf("Executing test => %v", tt.name)
		seed := int64(45)
		policy := ThompsonSampling{rewardValues: []float32{0, 1}, seed: &seed}
		armSelected := policy.SelectArm(tt.arms, tt.armsRewards)
		assert.Equal(t, tt.selection, armSelected)
	}
}
