package service

import (
	"context"
	"errors"
	"testing"

	"github.com/mattvictorino/rating-api/internal/domain"
	"github.com/mattvictorino/rating-api/internal/ports"
	"github.com/stretchr/testify/assert"
)

func TestGetByID(t *testing.T) {
	type TestGetByIDSetup struct {
		repository *ports.RatingRepositoryMock
		ctx        context.Context
		rating     domain.Rating
	}

	type TestGetByIDStruct struct {
		setup   TestGetByIDSetup
		asserts func(t *testing.T, rating domain.Rating, err error, env TestGetByIDSetup)
	}

	testsCases := map[string]TestGetByIDStruct{
		"returns rating when found": {
			setup: func() TestGetByIDSetup {
				repo := new(ports.RatingRepositoryMock)
				ctx := context.Background()
				rating := domain.Rating{
					ID:        "rating-123",
					UserID:    "user-123",
					ProductID: "product-123",
					Type:      "like",
					Value:     1,
				}

				repo.On("GetByID", ctx, rating.ID).Return(rating, nil)

				return TestGetByIDSetup{
					repository: repo,
					ctx:        ctx,
					rating:     rating,
				}
			}(),
			asserts: func(t *testing.T, rating domain.Rating, err error, env TestGetByIDSetup) {
				assert.Equal(t, env.rating, rating)
				assert.NoError(t, err)
			},
		},
		"returns no rating and error when not found": {
			setup: func() TestGetByIDSetup {
				repo := new(ports.RatingRepositoryMock)
				ctx := context.Background()
				id := "rating-123"

				repo.On("GetByID", ctx, id).Return(domain.Rating{}, domain.ErrNotFound)

				return TestGetByIDSetup{
					repository: repo,
					ctx:        ctx,
					rating: domain.Rating{
						ID: id,
					},
				}
			}(),
			asserts: func(t *testing.T, rating domain.Rating, err error, env TestGetByIDSetup) {
				assert.Equal(t, err, domain.ErrNotFound)
			},
		},
		"returns no rating and error when repository breaks": {
			setup: func() TestGetByIDSetup {
				repo := new(ports.RatingRepositoryMock)
				ctx := context.Background()
				id := "rating-123"

				repo.On("GetByID", ctx, id).Return(domain.Rating{}, errors.New("some error"))

				return TestGetByIDSetup{
					repository: repo,
					ctx:        ctx,
					rating: domain.Rating{
						ID: id,
					},
				}
			}(),
			asserts: func(t *testing.T, rating domain.Rating, err error, env TestGetByIDSetup) {
				assert.Error(t, err)
			},
		},
	}

	for name, tc := range testsCases {
		t.Run(name, func(t *testing.T) {
			env := tc.setup

			service := NewRatingService(env.repository)
			rating, err := service.GetByID(env.ctx, env.rating.ID)

			tc.asserts(t, rating, err, env)
		})
	}
}

func TestInsert(t *testing.T) {
	type TestInsertSetup struct {
		repository *ports.RatingRepositoryMock
		ctx        context.Context
		rating     domain.Rating
	}

	type TestInsertStruct struct {
		setup   TestInsertSetup
		asserts func(t *testing.T, err error, env TestInsertSetup)
	}

	testsCases := map[string]TestInsertStruct{
		"inserts rating successfully": {
			setup: func() TestInsertSetup {
				repo := new(ports.RatingRepositoryMock)
				ctx := context.Background()
				rating := domain.Rating{
					ID:        "rating-123",
					UserID:    "user-123",
					ProductID: "product-123",
					Type:      "like",
					Value:     1,
				}

				repo.On("Insert", ctx, rating).Return(nil)
				return TestInsertSetup{
					repository: repo,
					ctx:        ctx,
					rating:     rating,
				}
			}(),
			asserts: func(t *testing.T, err error, env TestInsertSetup) {
				assert.NoError(t, err)
			},
		},
		"returns error when repository breaks": {
			setup: func() TestInsertSetup {
				repo := new(ports.RatingRepositoryMock)
				ctx := context.Background()
				rating := domain.Rating{
					ID:        "rating-123",
					UserID:    "user-123",
					ProductID: "product-123",
					Type:      "like",
					Value:     1,
				}

				repo.On("Insert", ctx, rating).Return(errors.New("some error"))
				return TestInsertSetup{
					repository: repo,
					ctx:        ctx,
					rating:     rating,
				}
			}(),
			asserts: func(t *testing.T, err error, env TestInsertSetup) {
				assert.Error(t, err)
			},
		},
	}

	for name, tc := range testsCases {
		t.Run(name, func(t *testing.T) {
			env := tc.setup

			service := NewRatingService(env.repository)
			err := service.Insert(env.ctx, env.rating)

			tc.asserts(t, err, env)
		})
	}
}

func TestUpdate(t *testing.T) {
	type TestUpdateSetup struct {
		repository *ports.RatingRepositoryMock
		ctx        context.Context
		rating     domain.Rating
	}

	type TestUpdateStruct struct {
		setup   TestUpdateSetup
		asserts func(t *testing.T, err error, env TestUpdateSetup)
	}

	testsCases := map[string]TestUpdateStruct{
		"updates rating successfully": {
			setup: func() TestUpdateSetup {
				repo := new(ports.RatingRepositoryMock)
				ctx := context.Background()
				rating := domain.Rating{
					ID:        "rating-123",
					UserID:    "user-123",
					ProductID: "product-123",
					Type:      "like",
					Value:     1,
				}

				repo.On("Update", ctx, rating).Return(nil)
				return TestUpdateSetup{
					repository: repo,
					ctx:        ctx,
					rating:     rating,
				}
			}(),
			asserts: func(t *testing.T, err error, env TestUpdateSetup) {
				assert.NoError(t, err)
			},
		},
		"returns error when repository breaks": {
			setup: func() TestUpdateSetup {
				repo := new(ports.RatingRepositoryMock)
				ctx := context.Background()
				rating := domain.Rating{
					ID:        "rating-123",
					UserID:    "user-123",
					ProductID: "product-123",
					Type:      "like",
					Value:     1,
				}

				repo.On("Update", ctx, rating).Return(errors.New("some error"))
				return TestUpdateSetup{
					repository: repo,
					ctx:        ctx,
					rating:     rating,
				}
			}(),
			asserts: func(t *testing.T, err error, env TestUpdateSetup) {
				assert.Error(t, err)
			},
		},
	}

	for name, tc := range testsCases {
		t.Run(name, func(t *testing.T) {
			env := tc.setup

			service := NewRatingService(env.repository)
			err := service.Update(env.ctx, env.rating)

			tc.asserts(t, err, env)
		})
	}
}

func TestDelete(t *testing.T) {
	type TestDeleteSetup struct {
		repository *ports.RatingRepositoryMock
		ctx        context.Context
		ratingID   string
	}

	type TestDeleteStruct struct {
		setup   TestDeleteSetup
		asserts func(t *testing.T, err error, env TestDeleteSetup)
	}

	testsCases := map[string]TestDeleteStruct{
		"deletes rating successfully": {
			setup: func() TestDeleteSetup {
				repo := new(ports.RatingRepositoryMock)
				ctx := context.Background()
				ratingID := "rating-123"

				repo.On("Delete", ctx, ratingID).Return(nil)
				return TestDeleteSetup{
					repository: repo,
					ctx:        ctx,
					ratingID:   ratingID,
				}
			}(),
			asserts: func(t *testing.T, err error, env TestDeleteSetup) {
				assert.NoError(t, err)
			},
		},
		"returns error when repository breaks": {
			setup: func() TestDeleteSetup {
				repo := new(ports.RatingRepositoryMock)
				ctx := context.Background()
				ratingID := "rating-123"

				repo.On("Delete", ctx, ratingID).Return(errors.New("some error"))
				return TestDeleteSetup{
					repository: repo,
					ctx:        ctx,
					ratingID:   ratingID,
				}
			}(),
			asserts: func(t *testing.T, err error, env TestDeleteSetup) {
				assert.Error(t, err)
			},
		},
	}

	for name, tc := range testsCases {
		t.Run(name, func(t *testing.T) {
			env := tc.setup

			service := NewRatingService(env.repository)
			err := service.Delete(env.ctx, env.ratingID)

			tc.asserts(t, err, env)
		})
	}
}
