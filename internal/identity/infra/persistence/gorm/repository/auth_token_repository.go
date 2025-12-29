package repository

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/model"
	"github.com/cristiano-pacheco/goflix/internal/identity/domain/repository"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/persistence/gorm/entity"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/persistence/gorm/mapper"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/database"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/errs"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/otel"
)

type AuthTokenRepository struct {
	db     *database.GoflixDB
	mapper mapper.AuthTokenMapper
}

var _ repository.AuthTokenRepositoryI = (*AuthTokenRepository)(nil)

func NewAuthTokenRepository(db *database.GoflixDB, mapper mapper.AuthTokenMapper) *AuthTokenRepository {
	return &AuthTokenRepository{db, mapper}
}

func (r *AuthTokenRepository) Create(
	ctx context.Context,
	authTokenModel model.AuthTokenModel,
) (model.AuthTokenModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "AuthTokenRepository.Create")
	defer span.End()

	authTokenEntity := r.mapper.ToEntity(authTokenModel)
	result := r.db.WithContext(ctx).Create(&authTokenEntity)
	if result.Error != nil {
		return model.AuthTokenModel{}, result.Error
	}

	authTokenModel, err := r.mapper.ToModel(authTokenEntity)
	if err != nil {
		return model.AuthTokenModel{}, err
	}

	return authTokenModel, nil
}

func (r *AuthTokenRepository) Update(ctx context.Context, authTokenModel model.AuthTokenModel) error {
	ctx, span := otel.Trace().StartSpan(ctx, "AuthTokenRepository.Update")
	defer span.End()

	authTokenEntity := r.mapper.ToEntity(authTokenModel)
	result := r.db.WithContext(ctx).Save(&authTokenEntity)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *AuthTokenRepository) Delete(ctx context.Context, id uint64) error {
	ctx, span := otel.Trace().StartSpan(ctx, "AuthTokenRepository.Delete")
	defer span.End()

	result := r.db.WithContext(ctx).Delete(&entity.AuthTokenEntity{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errs.ErrNotFound
	}

	return nil
}

func (r *AuthTokenRepository) FindByToken(ctx context.Context, token string) (model.AuthTokenModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "AuthTokenRepository.FindByToken")
	defer span.End()

	var authTokenEntity entity.AuthTokenEntity
	result := r.db.WithContext(ctx).Where("token = ?", token).First(&authTokenEntity)

	if result.Error != nil {
		return model.AuthTokenModel{}, result.Error
	}

	if authTokenEntity.ID == 0 {
		return model.AuthTokenModel{}, errs.ErrNotFound
	}

	authTokenModel, err := r.mapper.ToModel(authTokenEntity)
	if err != nil {
		return model.AuthTokenModel{}, err
	}

	return authTokenModel, nil
}
