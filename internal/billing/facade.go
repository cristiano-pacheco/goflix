package billing

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/billing/ports"
)

type FacadeI interface {
	IsUserSubscriptionActive(ctx context.Context, userID uint64) (bool, error)
}

type Facade struct {
	subscriptionRepository ports.SubscriptionRepositoryI
}

func NewFacade(subscriptionRepository ports.SubscriptionRepositoryI) *Facade {
	return &Facade{
		subscriptionRepository,
	}
}

func (f *Facade) IsUserSubscriptionActive(ctx context.Context, userID uint64) (bool, error) {
	_, err := f.subscriptionRepository.FindActiveSubscriptionByUserID(ctx, userID)
	if err != nil {
		return false, err
	}

	return true, nil
}
