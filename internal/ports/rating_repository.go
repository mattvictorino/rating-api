package ports

import (
	"context"

	"github.com/mattvictorino/rating-api/internal/domain"
)

type RatingRepository interface {
	GetByID(ctx context.Context, id string) (domain.Rating, error)
	Insert(ctx context.Context, rating domain.Rating) error
	Update(ctx context.Context, rating domain.Rating) error
	Delete(ctx context.Context, id string) error
}
