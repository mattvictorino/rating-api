package service

import (
	"context"

	"github.com/mattvictorino/rating-api/internal/domain"
	"github.com/mattvictorino/rating-api/internal/ports"
)

type Rating struct {
	repository ports.RatingRepository
}

func NewRatingService(repository ports.RatingRepository) Rating {
	return Rating{
		repository: repository,
	}
}

func (s Rating) GetByID(ctx context.Context, id string) (domain.Rating, error) {
	return s.repository.GetByID(ctx, id)
}

func (s Rating) Insert(ctx context.Context, rating domain.Rating) error {
	return s.repository.Insert(ctx, rating)
}

func (s Rating) Update(ctx context.Context, rating domain.Rating) error {
	return s.repository.Update(ctx, rating)
}

func (s Rating) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}
