package entity

import "time"

type UserEntity struct {
	ID                     uint64     `gorm:"primarykey;autoIncrement;column:id"`
	Name                   string     `gorm:"type:varchar;not null;column:name"`
	Email                  string     `gorm:"type:varchar;not null;unique;column:email"`
	PasswordHash           string     `gorm:"type:varchar;not null;column:password_hash"`
	IsActivated            bool       `gorm:"type:boolean;not null;default:false;column:is_activated"`
	ConfirmationToken      *string    `gorm:"type:varchar;column:confirmation_token"`
	ConfirmationExpiresAt  *time.Time `gorm:"type:timestamptz;column:confirmation_expires_at"`
	ConfirmedAt            *time.Time `gorm:"type:timestamptz;column:confirmed_at"`
	ResetPasswordToken     *string    `gorm:"type:varchar;column:reset_password_token"`
	ResetPasswordExpiresAt *time.Time `gorm:"type:timestamptz;column:reset_password_expires_at"`
	CreatedAt              time.Time  `gorm:"type:timestamptz;default:now();column:created_at"`
	UpdatedAt              time.Time  `gorm:"type:timestamptz;default:now();column:updated_at"`
}

func (*UserEntity) TableName() string {
	return "users"
}
