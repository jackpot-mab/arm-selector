package controller

import "jackpot-mab/arm-selector/policy"

type Response struct {
	DecisionId string     `json:"decision_id"`
	Arm        policy.Arm `json:"arms"`
}
