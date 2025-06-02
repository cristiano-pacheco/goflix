package entity

import "time"

type AuthTokenEntity struct {
	ID        uint64    `gorm:"primarykey;autoIncrement;column:id"`
	UserID    uint64    `gorm:"type:bigint;not null;column:user_id"`
	Token     string    `gorm:"type:varchar;not null;unique;column:token"`
	ExpiresAt time.Time `gorm:"type:timestamptz;not null;column:expires_at"`
	CreatedAt time.Time `gorm:"type:timestamptz;default:now();column:created_at"`
	UpdatedAt time.Time `gorm:"type:timestamptz;default:now();column:updated_at"`
}

func (*AuthTokenEntity) TableName() string {
	return "login_token"
}
