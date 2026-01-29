package ports

import "github.com/mattvictorino/rating-api/internal/domain"

type Rating interface {
	GetByID(id string) (domain.Rating, error)
	Insert(rating domain.Rating) error
	Update(rating domain.Rating) error
	Delete(id string) error
}
