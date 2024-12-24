package model

import "time"

type AssignedOrder struct {
	AssignedOrderId             string    `json:"assigned_order_id"`
	ExecutorRating              float32   `json:"executor_rating"`
	OrderId                     string    `json:"order_id"`
	ExecutorId                  string    `json:"executor_id"`
	ExecutionStatus             string    `json:"execution_status"`
	CoinCoefficient             float32   `json:"coin_coefficient"`
	CoinBonusAmount             int32     `json:"coin_bonus_amount"`
	FinalCoinAmount             int32     `json:"final_coin_amount"`
	ZoneId                      string    `json:"zone_id"`
	HasExecutorFallbackBeenUsed bool      `json:"has_executor_fallback_been_used"`
	AssignTime                  time.Time `json:"assign_time"`
	FirstAcquireTime            time.Time `json:"first_acquire_time"`
	CompletedTime               time.Time `json:"completed_time"`
}

type DWHOrder struct {
	AssignTime      time.Time `json:"assign_time"`
	FinalCoinAmount int32     `json:"final_coin_amount"`
	ExecutionStatus string    `json:"execution_status"`
	AcquireSeconds  int32     `json:"acquire_seconds"`
	ExecutorRating  float32   `json:"executor_rating"`
	OrderId         string    `json:"order_id"`
	ExecutorId      string    `json:"executor_id"`
	ZoneId          string    `json:"zone_id"`
	CoinBonusAmount int32     `json:"coin_bonus_amount"`
}
