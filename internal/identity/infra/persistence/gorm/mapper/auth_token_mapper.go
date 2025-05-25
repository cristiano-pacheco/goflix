package mapper

import (
	"github.com/cristiano-pacheco/goflix/internal/identity/domain/model"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/persistence/gorm/entity"
)

type AuthTokenMapper interface {
	ToModel(entity entity.AuthTokenEntity) (model.AuthTokenModel, error)
	ToEntity(model model.AuthTokenModel) entity.AuthTokenEntity
}

type authTokenMapper struct {
}

func NewAuthTokenMapper() AuthTokenMapper {
	return &authTokenMapper{}
}

func (m *authTokenMapper) ToModel(entity entity.AuthTokenEntity) (model.AuthTokenModel, error) {
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

func (m *authTokenMapper) ToEntity(model model.AuthTokenModel) entity.AuthTokenEntity {
	return entity.AuthTokenEntity{
		ID:        model.ID(),
		UserID:    model.UserID(),
		Token:     model.Token(),
		ExpiresAt: model.ExpiresAt(),
		CreatedAt: model.CreatedAt(),
		UpdatedAt: model.UpdatedAt(),
	}
}
