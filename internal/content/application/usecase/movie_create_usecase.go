package usecase

import (
	"context"

	"github.com/cristiano-pacheco/goflix/internal/content/domain/model"
	"github.com/cristiano-pacheco/goflix/internal/content/ports"
)

type MovieCreateUsecase struct {
	repository ports.MovieRepositoryI
}

func NewMovieCreateUsecase(repository ports.MovieRepositoryI) *MovieCreateUsecase {
	return &MovieCreateUsecase{repository: repository}
}

type MovieCreateUsecaseInput struct {
	ExternalRating float64
	Content        model.ContentModel
	Video          model.VideoModel
	Thumbnail      model.ThumbnailModel
}

func (uc *MovieCreateUsecase) Execute(ctx context.Context, movie model.MovieModel) (model.MovieModel, error) {
	return uc.repository.Create(ctx, movie)
}
