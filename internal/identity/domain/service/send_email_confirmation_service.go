package service

import "context"

type SendEmailConfirmationService interface {
	Execute(ctx context.Context, userID uint64) error
}
