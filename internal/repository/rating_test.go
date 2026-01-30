package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/mattvictorino/rating-api/internal/database"
	"github.com/mattvictorino/rating-api/internal/domain"
	"github.com/mattvictorino/rating-api/internal/repository/persistence"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestRatingRepository_GetByID(t *testing.T) {
	type setupEnv struct {
		repo Rating
		ctx  context.Context
		id   string
	}

	type testCase struct {
		setup   func() setupEnv
		asserts func(t *testing.T, rating domain.Rating, err error)
	}

	validID := primitive.NewObjectID()
	validHex := validID.Hex()

	tests := map[string]testCase{

		"returns rating successfully": {
			setup: func() setupEnv {
				collection := new(database.MongoCollectionMock)
				ctx := context.Background()

				ratingDB := persistence.Rating{ID: validID}

				sr := mongo.NewSingleResultFromDocument(ratingDB, nil, nil)

				collection.On("FindOne", ctx, bson.M{"_id": validID}).Return(sr)

				repo := Rating{collection: collection}
				return setupEnv{repo, ctx, validHex}
			},
			asserts: func(t *testing.T, rating domain.Rating, err error) {
				assert.NoError(t, err)
				assert.NotEmpty(t, rating)
			},
		},

		"returns not found": {
			setup: func() setupEnv {
				collection := new(database.MongoCollectionMock)
				ctx := context.Background()

				var rating persistence.Rating
				sr := mongo.NewSingleResultFromDocument(&rating, mongo.ErrNoDocuments, nil)
				collection.On("FindOne", ctx, bson.M{"_id": validID}).Return(sr)

				repo := Rating{collection: collection}
				return setupEnv{repo, ctx, validHex}
			},
			asserts: func(t *testing.T, rating domain.Rating, err error) {
				assert.ErrorIs(t, err, domain.ErrNotFound)
			},
		},

		"invalid id format": {
			setup: func() setupEnv {
				return setupEnv{repo: Rating{}, ctx: context.Background(), id: "invalid"}
			},
			asserts: func(t *testing.T, rating domain.Rating, err error) {
				assert.Error(t, err)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			env := tc.setup()
			r, err := env.repo.GetByID(env.ctx, env.id)
			tc.asserts(t, r, err)
		})
	}
}

func TestRatingRepository_Insert(t *testing.T) {
	type env struct {
		repo   Rating
		ctx    context.Context
		rating domain.Rating
	}

	tests := map[string]struct {
		setup   func() env
		asserts func(t *testing.T, err error)
	}{

		"inserts successfully": {
			setup: func() env {
				collection := new(database.MongoCollectionMock)
				ctx := context.Background()
				rating := domain.Rating{ProductID: "p1"}

				collection.On("InsertOne", ctx, mock.Anything).
					Return(&mongo.InsertOneResult{}, nil)

				return env{Rating{collection}, ctx, rating}
			},
			asserts: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},

		"mongo returns error": {
			setup: func() env {
				collection := new(database.MongoCollectionMock)
				ctx := context.Background()

				collection.On("InsertOne", ctx, mock.Anything).
					Return((*mongo.InsertOneResult)(nil), errors.New("db error"))

				return env{Rating{collection}, ctx, domain.Rating{}}
			},
			asserts: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			env := tc.setup()
			err := env.repo.Insert(env.ctx, env.rating)
			tc.asserts(t, err)
		})
	}
}

func TestRatingRepository_Update(t *testing.T) {
	validID := primitive.NewObjectID()

	tests := map[string]struct {
		setup   func() (Rating, context.Context, domain.Rating)
		asserts func(t *testing.T, err error)
	}{

		"updates successfully": {
			setup: func() (Rating, context.Context, domain.Rating) {
				collection := new(database.MongoCollectionMock)
				ctx := context.Background()

				collection.On("UpdateOne", ctx, mock.Anything, mock.Anything).
					Return(&mongo.UpdateResult{MatchedCount: 1}, nil)

				return Rating{collection}, ctx, domain.Rating{ID: validID.Hex()}
			},
			asserts: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},

		"not found": {
			setup: func() (Rating, context.Context, domain.Rating) {
				collection := new(database.MongoCollectionMock)
				ctx := context.Background()

				collection.On("UpdateOne", ctx, mock.Anything, mock.Anything).
					Return(&mongo.UpdateResult{MatchedCount: 0}, nil)

				return Rating{collection}, ctx, domain.Rating{ID: validID.Hex()}
			},
			asserts: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, domain.ErrNotFound)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			repo, ctx, rating := tc.setup()
			err := repo.Update(ctx, rating)
			tc.asserts(t, err)
		})
	}
}

func TestRatingRepository_Delete(t *testing.T) {
	validID := primitive.NewObjectID()

	tests := map[string]struct {
		setup   func() (Rating, context.Context, string)
		asserts func(t *testing.T, err error)
	}{

		"deletes successfully": {
			setup: func() (Rating, context.Context, string) {
				collection := new(database.MongoCollectionMock)
				ctx := context.Background()

				collection.On("DeleteOne", ctx, mock.Anything).
					Return(&mongo.DeleteResult{}, nil)

				return Rating{collection}, ctx, validID.Hex()
			},
			asserts: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			repo, ctx, id := tc.setup()
			err := repo.Delete(ctx, id)
			tc.asserts(t, err)
		})
	}
}
