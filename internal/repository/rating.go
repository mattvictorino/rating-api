package repository

import (
	"context"
	"errors"
	"time"

	"github.com/mattvictorino/rating-api/internal/database"
	"github.com/mattvictorino/rating-api/internal/domain"
	"github.com/mattvictorino/rating-api/internal/mapper"
	"github.com/mattvictorino/rating-api/internal/repository/persistence"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Rating struct {
	collection database.MongoCollection
}

func NewRatingRepository(mongo *database.Mongo) Rating {
	return Rating{
		collection: mongo.DB.Collection("ratings"),
	}
}

func (r Rating) GetByID(ctx context.Context, id string) (domain.Rating, error) {
	objID, err := idtoObjectID(id)
	if err != nil {
		return domain.Rating{}, err
	}

	filter := bson.M{"_id": objID}

	var ratingDB persistence.Rating
	err = r.collection.FindOne(ctx, filter).Decode(&ratingDB)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return domain.Rating{}, domain.ErrNotFound
		}

		return domain.Rating{}, err
	}

	rating := mapper.ToRatingDomain(ratingDB)
	return rating, nil
}

func (r Rating) Insert(ctx context.Context, rating domain.Rating) error {
	ratingDB, err := mapper.ToRatingDB(rating)
	if err == nil {
		now := time.Now()

		ratingDB.InsertedAt = now
		ratingDB.UpdatedAt = now

		_, err = r.collection.InsertOne(ctx, ratingDB)
	}

	return err
}

func (r Rating) Update(ctx context.Context, rating domain.Rating) error {
	objID, err := idtoObjectID(rating.ID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	update := bson.M{
		"$set": bson.M{
			"product_id": rating.ProductID,
			"type":       rating.Type,
			"value":      rating.Value,
			"updated_at": time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return domain.ErrNotFound
	}

	return nil
}

func (r Rating) Delete(ctx context.Context, id string) error {
	objID, err := idtoObjectID(id)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}

	_, err = r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func idtoObjectID(id string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NewObjectID(), err
	}

	return objID, nil
}
