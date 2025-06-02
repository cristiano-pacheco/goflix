package entity

import "time"

type SubscriptionEntity struct {
	ID        uint64    `gorm:"primarykey;autoIncrement;column:id"`
	UserID    uint64    `gorm:"type:bigint;not null;column:user_id"`
	PlanID    uint64    `gorm:"type:bigint;not null;column:plan_id"`
	Status    string    `gorm:"type:varchar(20);not null;column:status"`
	StartDate time.Time `gorm:"type:timestamptz;not null;column:start_date"`
	EndDate   time.Time `gorm:"type:timestamptz;column:end_date"`
	AutoRenew bool      `gorm:"type:boolean;not null;default:true;column:auto_renew"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:now();column:created_at"`
	UpdatedAt time.Time `gorm:"type:timestamptz;default:now();column:updated_at"`
}

func (*SubscriptionEntity) TableName() string {
	return "subscription"
}
