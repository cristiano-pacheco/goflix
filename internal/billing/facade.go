package billing

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/billing/ports"
)

type BillingFacadeI interface {
	IsUserSubscriptionActive(ctx context.Context, userID uint64) (bool, error)
}

type BillingFacade struct {
	subscriptionRepository ports.SubscriptionRepositoryI
}

func NewBillingFacade(subscriptionRepository ports.SubscriptionRepositoryI) *BillingFacade {
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
