package billing

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/billing/domain/repository"
)

type BillingFacadeI interface {
	IsUserSubscriptionActive(ctx context.Context, userID uint64) (bool, error)
}

type BillingFacade struct {
	subscriptionRepository repository.SubscriptionRepositoryI
}

func NewBillingFacade(subscriptionRepository repository.SubscriptionRepositoryI) *BillingFacade {
	return &BillingFacade{
		subscriptionRepository,
	}
}

func (f *BillingFacade) IsUserSubscriptionActive(ctx context.Context, userID uint64) (bool, error) {
	_, err := f.subscriptionRepository.FindActiveSubscriptionByUserID(ctx, userID)
	if err != nil {
		return false, err
	}

	return true, nil
}
