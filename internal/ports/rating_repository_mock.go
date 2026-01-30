package ports

import (
	"context"

	"github.com/mattvictorino/rating-api/internal/domain"
	"github.com/stretchr/testify/mock"
)

type RatingRepositoryMock struct {
	mock.Mock
}

func (m *RatingRepositoryMock) GetByID(ctx context.Context, id string) (domain.Rating, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(domain.Rating), args.Error(1)
}

func (m *RatingRepositoryMock) Insert(ctx context.Context, rating domain.Rating) error {
	args := m.Called(ctx, rating)
	return args.Error(0)
}

func (m *RatingRepositoryMock) Update(ctx context.Context, rating domain.Rating) error {
	args := m.Called(ctx, rating)
	return args.Error(0)
}

func (m *RatingRepositoryMock) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
