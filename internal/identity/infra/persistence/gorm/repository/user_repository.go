package repository

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/identity/domain/model"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/persistence/gorm/entity"
	"github.com/cristiano-pacheco/goflix/internal/identity/infra/persistence/gorm/mapper"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/database"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/errs"
	"github.com/cristiano-pacheco/goflix/internal/shared/modules/otel"
)

type UserRepository struct {
	db     *database.GoflixDB
	mapper mapper.UserMapper
}

func NewUserRepository(db *database.GoflixDB, mapper mapper.UserMapper) *UserRepository {
	return &UserRepository{db, mapper}
}

func (r *UserRepository) Create(ctx context.Context, userModel model.UserModel) (model.UserModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "UserRepository.Create")
	defer span.End()

	userEntity := r.mapper.ToEntity(userModel)
	result := r.db.WithContext(ctx).Create(&userEntity)
	if result.Error != nil {
		return model.UserModel{}, result.Error
	}
	userModel, err := r.mapper.ToModel(userEntity)
	if err != nil {
		return model.UserModel{}, err
	}
	return userModel, nil
}

func (r *UserRepository) Update(ctx context.Context, model model.UserModel) error {
	ctx, span := otel.Trace().StartSpan(ctx, "UserRepository.Update")
	defer span.End()

	userEntity := r.mapper.ToEntity(model)
	result := r.db.WithContext(ctx).Save(&userEntity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id uint64) (model.UserModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "UserRepository.FindByID")
	defer span.End()

	var userEntity entity.UserEntity
	r.db.WithContext(ctx).First(&userEntity, id)
	if userEntity.ID == 0 {
		return model.UserModel{}, errs.ErrNotFound
	}

	userModel, err := r.mapper.ToModel(userEntity)
	if err != nil {
		return model.UserModel{}, err
	}

	return userModel, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (model.UserModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "UserRepository.FindByEmail")
	defer span.End()
	var userEntity entity.UserEntity
	r.db.WithContext(ctx).Where("email = ?", email).First(&userEntity)
	if userEntity.ID == 0 {
		return model.UserModel{}, errs.ErrNotFound
	}
	userModel, err := r.mapper.ToModel(userEntity)
	if err != nil {
		return model.UserModel{}, err
	}
	return userModel, nil
}

func (r *UserRepository) FindByConfirmationToken(ctx context.Context, token string) (model.UserModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "UserRepository.FindByConfirmationToken")
	defer span.End()
	var userEntity entity.UserEntity
	r.db.WithContext(ctx).Where("confirmation_token = ?", token).First(&userEntity)
	if userEntity.ID == 0 {
		return model.UserModel{}, errs.ErrNotFound
	}
	userModel, err := r.mapper.ToModel(userEntity)
	if err != nil {
		return model.UserModel{}, err
	}
	return userModel, nil
}

func (r *UserRepository) FindByResetPasswordToken(ctx context.Context, token string) (model.UserModel, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "UserRepository.FindByResetPasswordToken")
	defer span.End()
	var userEntity entity.UserEntity
	r.db.WithContext(ctx).Where("reset_password_token = ?", token).First(&userEntity)
	if userEntity.ID == 0 {
		return model.UserModel{}, errs.ErrNotFound
	}
	userModel, err := r.mapper.ToModel(userEntity)
	if err != nil {
		return model.UserModel{}, err
	}
	return userModel, nil
}

func (r *UserRepository) IsActivated(ctx context.Context, userID uint64) (bool, error) {
	ctx, span := otel.Trace().StartSpan(ctx, "UserRepository.IsActivated")
	defer span.End()
	var userEntity entity.UserEntity
	r.db.WithContext(ctx).Where("id = ?", userID).First(&userEntity)
	if userEntity.ID == 0 {
		return false, errs.ErrNotFound
	}
	return userEntity.IsActivated, nil
}
