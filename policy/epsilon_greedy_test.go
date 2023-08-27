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
