package entity

import "time"

type PlanEntity struct {
	ID          uint64    `gorm:"primarykey;autoIncrement;column:id"`
	Name        string    `gorm:"type:varchar(100);not null;column:name"`
	Description string    `gorm:"type:varchar(1000);not null;column:description"`
	AmountCents uint      `gorm:"type:integer;not null;column:amount_cents"`
	Currency    string    `gorm:"type:varchar(3);not null;column:currency"`
	Interval    string    `gorm:"type:varchar(10);not null;column:interval"`
	TrialPeriod uint      `gorm:"type:integer;not null;column:trial_period"`
	CreatedAt   time.Time `gorm:"type:timestamptz;default:now();column:created_at"`
	UpdatedAt   time.Time `gorm:"type:timestamptz;default:now();column:updated_at"`
}

func (*PlanEntity) TableName() string {
	return "plan"
}
