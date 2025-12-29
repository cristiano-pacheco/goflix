package service

import "context"

type SendEmailConfirmationServiceI interface {
	Execute(ctx context.Context, userID uint64) error
}
