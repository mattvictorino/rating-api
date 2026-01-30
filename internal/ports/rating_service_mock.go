package ports

import (
	"context"

	"github.com/mattvictorino/rating-api/internal/domain"
	"github.com/stretchr/testify/mock"
)

type RatingServiceMock struct {
	mock.Mock
}

func (r *RatingServiceMock) GetByID(ctx context.Context, ID string) (domain.Rating, error) {
	args := r.Called(ctx, ID)
	return args.Get(0).(domain.Rating), args.Error(1)
}

func (r *RatingServiceMock) Insert(ctx context.Context, rating domain.Rating) error {
	args := r.Called(ctx, rating)
	return args.Error(0)
}

func (r *RatingServiceMock) Update(ctx context.Context, rating domain.Rating) error {
	args := r.Called(ctx, rating)
	return args.Error(0)
}

func (r *RatingServiceMock) Delete(ctx context.Context, id string) error {
	args := r.Called(ctx, id)
	return args.Error(0)
}
