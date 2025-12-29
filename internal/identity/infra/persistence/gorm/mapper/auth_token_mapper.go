package mapper

import (
	"github.com/cristiano-pacheco/goflix/internal/identity/domain/model"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/persistence/gorm/entity"
)

type AuthTokenMapperI interface {
	ToModel(entity entity.AuthTokenEntity) (model.AuthTokenModel, error)
	ToEntity(model model.AuthTokenModel) entity.AuthTokenEntity
}

type AuthTokenMapper struct {
}

var _ AuthTokenMapperI = (*AuthTokenMapper)(nil)

func NewAuthTokenMapper() *AuthTokenMapper {
	return &AuthTokenMapper{}
}

func (m *AuthTokenMapper) ToModel(entity entity.AuthTokenEntity) (model.AuthTokenModel, error) {
	authTokenModel, err := model.RestoreAuthTokenModel(
		entity.ID,
		entity.UserID,
		entity.Token,
		entity.ExpiresAt,
		entity.CreatedAt,
		entity.UpdatedAt,
	)
	if err != nil {
		return model.AuthTokenModel{}, err
	}
	return authTokenModel, nil
}

func (m *AuthTokenMapper) ToEntity(model model.AuthTokenModel) entity.AuthTokenEntity {
	return entity.AuthTokenEntity{
		ID:        model.ID(),
		UserID:    model.UserID(),
		Token:     model.Token(),
		ExpiresAt: model.ExpiresAt(),
		CreatedAt: model.CreatedAt(),
		UpdatedAt: model.UpdatedAt(),
	}
}
