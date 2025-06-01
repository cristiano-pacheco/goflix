package billing

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/repository"
)

type FacadeInterface interface {
	IsUserSubscriptionActive(ctx context.Context, userID uint64) (bool, error)
}

type facade struct {
	subscriptionRepository repository.SubscriptionRepository
}

func NewFacade(subscriptionRepository repository.SubscriptionRepository) FacadeInterface {
	return &facade{
		subscriptionRepository,
	}
}

func (f *facade) IsUserSubscriptionActive(ctx context.Context, userID uint64) (bool, error) {
	_, err := f.subscriptionRepository.FindActiveSubscriptionByUserID(ctx, userID)
	if err != nil {
		return false, err
	}

	return true, nil
}
