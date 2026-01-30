package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mattvictorino/rating-api/api/validator"
	"github.com/mattvictorino/rating-api/internal/domain"
	"github.com/mattvictorino/rating-api/internal/ports"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	type TestGetSetup struct {
		ctx     echo.Context
		rec     *httptest.ResponseRecorder
		service *ports.RatingServiceMock
	}

	type TestGetStruct struct {
		setup   TestGetSetup
		asserts func(t *testing.T, setup TestGetSetup)
	}

	testsCases := map[string]TestGetStruct{
		"should return 200 OK when success": {
			setup: func() TestGetSetup {
				e := echo.New()

				req := httptest.NewRequest(http.MethodGet, "/ratings/1", nil)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)

				ctx.SetParamNames("id")
				ctx.SetParamValues("1")

				serviceMock := new(ports.RatingServiceMock)
				serviceMock.On("GetByID", ctx.Request().Context(), "1").Return(domain.Rating{}, nil)

				return TestGetSetup{
					ctx:     ctx,
					rec:     rec,
					service: serviceMock,
				}
			}(),
			asserts: func(t *testing.T, setup TestGetSetup) {
				assert.Equal(t, http.StatusOK, setup.ctx.Response().Status)
			},
		},
		"should return 404 not found when not getting a rating": {
			setup: func() TestGetSetup {
				e := echo.New()

				req := httptest.NewRequest(http.MethodGet, "/ratings/100", nil)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)

				ctx.SetParamNames("id")
				ctx.SetParamValues("100")

				serviceMock := new(ports.RatingServiceMock)
				serviceMock.On("GetByID", ctx.Request().Context(), "100").Return(domain.Rating{}, domain.ErrNotFound)

				return TestGetSetup{
					ctx:     ctx,
					rec:     rec,
					service: serviceMock,
				}
			}(),
			asserts: func(t *testing.T, setup TestGetSetup) {
				assert.Equal(t, http.StatusNotFound, setup.ctx.Response().Status)
			},
		},
		"should return 500 internal server error when service breaks": {
			setup: func() TestGetSetup {
				e := echo.New()

				req := httptest.NewRequest(http.MethodGet, "/ratings/100", nil)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)

				ctx.SetParamNames("id")
				ctx.SetParamValues("100")

				serviceMock := new(ports.RatingServiceMock)
				serviceMock.On("GetByID", ctx.Request().Context(), "100").Return(domain.Rating{}, errors.New("some error"))

				return TestGetSetup{
					ctx:     ctx,
					rec:     rec,
					service: serviceMock,
				}
			}(),
			asserts: func(t *testing.T, setup TestGetSetup) {
				assert.Equal(t, http.StatusInternalServerError, setup.ctx.Response().Status)
			},
		},
	}

	for name, tc := range testsCases {
		t.Run(name, func(t *testing.T) {
			env := tc.setup

			handler := NewHandler(env.service)
			err := handler.Get(env.ctx)

			require.NoError(t, err)
			tc.asserts(t, env)
		})
	}
}

func TestPost(t *testing.T) {
	type TestPostSetup struct {
		ctx     echo.Context
		rec     *httptest.ResponseRecorder
		service *ports.RatingServiceMock
		rating  domain.Rating
	}

	type TestPostStruct struct {
		setup   TestPostSetup
		asserts func(t *testing.T, setup TestPostSetup)
	}

	testsCases := map[string]TestPostStruct{
		"should return 201 created when success": {
			setup: func() TestPostSetup {
				e := echo.New()

				body := `{"product_id":"44ad1de9-a6ad-4f7e-b6fc-09ec9adcf81b", "type":"five_stars", "value": 1}`
				req := httptest.NewRequest(http.MethodPost, "/ratings", strings.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				e.Validator = validator.NewCustomValidator()

				ratingMock := domain.Rating{
					ProductID: "44ad1de9-a6ad-4f7e-b6fc-09ec9adcf81b",
					Type:      "five_stars",
					Value:     1,
				}

				serviceMock := new(ports.RatingServiceMock)
				serviceMock.On("Insert", ctx.Request().Context(), ratingMock).Return(nil)

				return TestPostSetup{
					ctx:     ctx,
					rec:     rec,
					service: serviceMock,
					rating:  ratingMock,
				}
			}(),
			asserts: func(t *testing.T, setup TestPostSetup) {
				assert.Equal(t, http.StatusCreated, setup.ctx.Response().Status)
			},
		},
		"should return 400 bad request when payload is incorrect": {
			setup: func() TestPostSetup {
				e := echo.New()

				body := `{"product_id":"", "type":"five_stars", "value": 1}`
				req := httptest.NewRequest(http.MethodPost, "/ratings", strings.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				e.Validator = validator.NewCustomValidator()

				ratingMock := domain.Rating{
					ProductID: "",
					Type:      "five_stars",
					Value:     1,
				}

				serviceMock := new(ports.RatingServiceMock)
				serviceMock.On("Insert", ctx.Request().Context(), ratingMock).Return(nil)

				return TestPostSetup{
					ctx:     ctx,
					rec:     rec,
					service: serviceMock,
					rating:  ratingMock,
				}
			}(),
			asserts: func(t *testing.T, setup TestPostSetup) {
				assert.Equal(t, http.StatusBadRequest, setup.ctx.Response().Status)
			},
		},
		"should return 500 internal server error when service breaks": {
			setup: func() TestPostSetup {
				e := echo.New()

				body := `{"product_id":"44ad1de9-a6ad-4f7e-b6fc-09ec9adcf81b", "type":"five_stars", "value": 1}`
				req := httptest.NewRequest(http.MethodPost, "/ratings", strings.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				e.Validator = validator.NewCustomValidator()

				ratingMock := domain.Rating{
					ProductID: "44ad1de9-a6ad-4f7e-b6fc-09ec9adcf81b",
					Type:      "five_stars",
					Value:     1,
				}

				serviceMock := new(ports.RatingServiceMock)
				serviceMock.On("Insert", ctx.Request().Context(), ratingMock).Return(errors.New("some error"))

				return TestPostSetup{
					ctx:     ctx,
					rec:     rec,
					service: serviceMock,
					rating:  ratingMock,
				}
			}(),
			asserts: func(t *testing.T, setup TestPostSetup) {
				assert.Equal(t, http.StatusInternalServerError, setup.ctx.Response().Status)
			},
		},
	}

	for name, tc := range testsCases {
		t.Run(name, func(t *testing.T) {
			env := tc.setup

			handler := NewHandler(env.service)
			err := handler.Post(env.ctx)

			require.NoError(t, err)
			tc.asserts(t, env)
		})
	}
}

func TestPut(t *testing.T) {
	type TestPutSetup struct {
		ctx     echo.Context
		rec     *httptest.ResponseRecorder
		service *ports.RatingServiceMock
		rating  domain.Rating
	}

	type TestPutStruct struct {
		setup   TestPutSetup
		asserts func(t *testing.T, setup TestPutSetup)
	}

	testsCases := map[string]TestPutStruct{
		"should return 200 OK when success": {
			setup: func() TestPutSetup {
				e := echo.New()
				e.Validator = validator.NewCustomValidator()

				body := `{"product_id":"44ad1de9-a6ad-4f7e-b6fc-09ec9adcf81b", "type":"five_stars", "value": 1}`
				req := httptest.NewRequest(http.MethodPut, "/ratings/1", strings.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.SetParamNames("id")
				ctx.SetParamValues("1")

				ratingMock := domain.Rating{
					ID:        "1",
					ProductID: "44ad1de9-a6ad-4f7e-b6fc-09ec9adcf81b",
					Type:      "five_stars",
					Value:     1,
				}

				serviceMock := new(ports.RatingServiceMock)
				serviceMock.On("Update", ctx.Request().Context(), ratingMock).Return(nil)

				return TestPutSetup{
					ctx:     ctx,
					rec:     rec,
					service: serviceMock,
					rating:  ratingMock,
				}
			}(),
			asserts: func(t *testing.T, setup TestPutSetup) {
				assert.Equal(t, http.StatusOK, setup.ctx.Response().Status)
			},
		},
		"should return 400 bad request when payload is incorrect": {
			setup: func() TestPutSetup {
				e := echo.New()
				e.Validator = validator.NewCustomValidator()

				body := `{"product_id":"", "type":"five_stars", "value": 1}`
				req := httptest.NewRequest(http.MethodPut, "/ratings/1", strings.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.SetParamNames("id")
				ctx.SetParamValues("1")

				ratingMock := domain.Rating{
					ID:        "1",
					ProductID: "",
					Type:      "five_stars",
					Value:     1,
				}

				serviceMock := new(ports.RatingServiceMock)
				serviceMock.On("Update", ctx.Request().Context(), ratingMock).Return(nil)

				return TestPutSetup{
					ctx:     ctx,
					rec:     rec,
					service: serviceMock,
					rating:  ratingMock,
				}
			}(),
			asserts: func(t *testing.T, setup TestPutSetup) {
				assert.Equal(t, http.StatusBadRequest, setup.ctx.Response().Status)
			},
		},
		"should return 404 not found when not getting a rating": {
			setup: func() TestPutSetup {
				e := echo.New()
				e.Validator = validator.NewCustomValidator()

				body := `{"product_id":"44ad1de9-a6ad-4f7e-b6fc-09ec9adcf81b", "type":"five_stars", "value": 1}`
				req := httptest.NewRequest(http.MethodPut, "/ratings/100", strings.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.SetParamNames("id")
				ctx.SetParamValues("100")

				ratingMock := domain.Rating{
					ID:        "100",
					ProductID: "44ad1de9-a6ad-4f7e-b6fc-09ec9adcf81b",
					Type:      "five_stars",
					Value:     1,
				}

				serviceMock := new(ports.RatingServiceMock)
				serviceMock.On("Update", ctx.Request().Context(), ratingMock).Return(domain.ErrNotFound)

				return TestPutSetup{
					ctx:     ctx,
					rec:     rec,
					service: serviceMock,
					rating:  ratingMock,
				}
			}(),
			asserts: func(t *testing.T, setup TestPutSetup) {
				assert.Equal(t, http.StatusNotFound, setup.ctx.Response().Status)
			},
		},
		"should return 500 internal server error when service breaks": {
			setup: func() TestPutSetup {
				e := echo.New()
				e.Validator = validator.NewCustomValidator()

				body := `{"product_id":"44ad1de9-a6ad-4f7e-b6fc-09ec9adcf81b", "type":"five_stars", "value": 1}`
				req := httptest.NewRequest(http.MethodPut, "/ratings/1", strings.NewReader(body))
				req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)
				ctx.SetParamNames("id")
				ctx.SetParamValues("1")

				ratingMock := domain.Rating{
					ID:        "1",
					ProductID: "44ad1de9-a6ad-4f7e-b6fc-09ec9adcf81b",
					Type:      "five_stars",
					Value:     1,
				}

				serviceMock := new(ports.RatingServiceMock)
				serviceMock.On("Update", ctx.Request().Context(), ratingMock).Return(errors.New("some error"))

				return TestPutSetup{
					ctx:     ctx,
					rec:     rec,
					service: serviceMock,
					rating:  ratingMock,
				}
			}(),
			asserts: func(t *testing.T, setup TestPutSetup) {
				assert.Equal(t, http.StatusInternalServerError, setup.ctx.Response().Status)
			},
		},
	}

	for name, tc := range testsCases {
		t.Run(name, func(t *testing.T) {
			env := tc.setup

			handler := NewHandler(env.service)
			err := handler.Put(env.ctx)

			require.NoError(t, err)
			tc.asserts(t, env)
		})
	}
}

func TestDelete(t *testing.T) {
	type TestDeleteSetup struct {
		ctx     echo.Context
		rec     *httptest.ResponseRecorder
		service *ports.RatingServiceMock
	}

	type TestDeleteStruct struct {
		setup   TestDeleteSetup
		asserts func(t *testing.T, setup TestDeleteSetup)
	}

	testsCases := map[string]TestDeleteStruct{
		"should return 200 OK when success": {
			setup: func() TestDeleteSetup {
				e := echo.New()

				req := httptest.NewRequest(http.MethodDelete, "/ratings/1", nil)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)

				ctx.SetParamNames("id")
				ctx.SetParamValues("1")

				serviceMock := new(ports.RatingServiceMock)
				serviceMock.On("Delete", ctx.Request().Context(), "1").Return(nil)

				return TestDeleteSetup{
					ctx:     ctx,
					rec:     rec,
					service: serviceMock,
				}
			}(),
			asserts: func(t *testing.T, setup TestDeleteSetup) {
				assert.Equal(t, http.StatusOK, setup.ctx.Response().Status)
			},
		},
		"should return 500 internal server error when service breaks": {
			setup: func() TestDeleteSetup {
				e := echo.New()

				req := httptest.NewRequest(http.MethodDelete, "/ratings/100", nil)
				rec := httptest.NewRecorder()
				ctx := e.NewContext(req, rec)

				ctx.SetParamNames("id")
				ctx.SetParamValues("100")

				serviceMock := new(ports.RatingServiceMock)
				serviceMock.On("Delete", ctx.Request().Context(), "100").Return(errors.New("some error"))

				return TestDeleteSetup{
					ctx:     ctx,
					rec:     rec,
					service: serviceMock,
				}
			}(),
			asserts: func(t *testing.T, setup TestDeleteSetup) {
				assert.Equal(t, http.StatusInternalServerError, setup.ctx.Response().Status)
			},
		},
	}

	for name, tc := range testsCases {
		t.Run(name, func(t *testing.T) {
			env := tc.setup

			handler := NewHandler(env.service)
			err := handler.Delete(env.ctx)

			require.NoError(t, err)
			tc.asserts(t, env)
		})
	}
}
