package ports

import (
	"github.com/mattvictorino/rating-api/internal/domain"
	"github.com/stretchr/testify/mock"
)

type RatingMock struct {
	mock.Mock
}

func (r *RatingMock) GetByID(ID string) (domain.Rating, error) {
	args := r.Called(ID)
	return args.Get(0).(domain.Rating), args.Error(1)
}

func (r *RatingMock) Insert(rating domain.Rating) error {
	args := r.Called(rating)
	return args.Error(0)
}

func (r *RatingMock) Update(rating domain.Rating) error {
	args := r.Called(rating)
	return args.Error(0)
}

func (r *RatingMock) Delete(id string) error {
	args := r.Called(id)
	return args.Error(0)
}
