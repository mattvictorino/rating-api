package service

import "github.com/mattvictorino/rating-api/internal/domain"

type Rating struct{}

func NewRatingService() Rating {
	return Rating{}
}

func (s Rating) GetByID(id string) (domain.Rating, error) {
	return domain.Rating{}, nil
}

func (s Rating) Insert(rating domain.Rating) error {
	return nil
}

func (s Rating) Update(rating domain.Rating) error {
	return nil
}

func (s Rating) Delete(id string) error {
	return nil
}
