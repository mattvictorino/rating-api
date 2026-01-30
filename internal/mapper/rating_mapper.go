package mapper

import (
	"github.com/mattvictorino/rating-api/internal/domain"
	"github.com/mattvictorino/rating-api/internal/repository/persistence"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToRatingDB(r domain.Rating) (persistence.Rating, error) {
	var objID primitive.ObjectID
	var err error

	if r.ID != "" {
		objID, err = primitive.ObjectIDFromHex(r.ID)
		if err != nil {
			return persistence.Rating{}, err
		}
	}

	return persistence.Rating{
		ID:         objID,
		UserID:     r.UserID,
		ProductID:  r.ProductID,
		Type:       r.Type,
		Value:      r.Value,
		InsertedAt: r.InsertedAt,
		UpdatedAt:  r.UpdatedAt,
	}, nil
}

func ToRatingDomain(r persistence.Rating) domain.Rating {
	return domain.Rating{
		ID:         r.ID.Hex(),
		UserID:     r.UserID,
		ProductID:  r.ProductID,
		Type:       r.Type,
		Value:      r.Value,
		InsertedAt: r.InsertedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}
