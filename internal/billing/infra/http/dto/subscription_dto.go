package dto

import "time"

type CreateSubscriptionRequest struct {
	PlanID uint64 `json:"plan_id"`
}

type CreateSubscriptionResponse struct {
	SubscriptionID uint64     `json:"subscription_id"`
	UserID         uint64     `json:"user_id"`
	PlanID         uint64     `json:"plan_id"`
	Status         string     `json:"status"`
	StartDate      time.Time  `json:"start_date"`
	EndDate        *time.Time `json:"end_date"`
	AutoRenew      bool       `json:"auto_renew"`
}

type SubscriptionResponse struct {
	SubscriptionID uint64     `json:"subscription_id"`
	UserID         uint64     `json:"user_id"`
	PlanID         uint64     `json:"plan_id"`
	Status         string     `json:"status"`
	StartDate      time.Time  `json:"start_date"`
	EndDate        *time.Time `json:"end_date"`
	AutoRenew      bool       `json:"auto_renew"`
}

type ListSubscriptionsResponse struct {
	Subscriptions []SubscriptionResponse `json:"subscriptions"`
}

type IsUserSubscriptionActiveResponse struct {
	IsActive bool `json:"is_active"`
}
