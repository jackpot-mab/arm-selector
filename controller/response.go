package controller

import "jackpot-mab/arm-selector/policy"

type Response struct {
	DecisionId string
	Arm        policy.Arm
}
