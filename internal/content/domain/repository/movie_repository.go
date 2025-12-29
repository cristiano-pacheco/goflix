package repository

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/content/domain/model"
)

type MovieRepository interface {
	Create(ctx context.Context, movie model.MovieModel) (model.MovieModel, error)
	Update(ctx context.Context, movie model.MovieModel) error
	Delete(ctx context.Context, id uint64) error
	FindByID(ctx context.Context, id uint64) (model.MovieModel, error)
	Find(ctx context.Context) ([]model.MovieModel, error)
}
