package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mattvictorino/rating-api/api/schema/request"
	"github.com/mattvictorino/rating-api/internal/domain"
	"github.com/mattvictorino/rating-api/internal/ports"
)

type Handler struct {
	service ports.Rating
}

func NewHandler(service ports.Rating) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) Get(c echo.Context) error {
	rating, err := h.service.GetByID(c.Param("id"))
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return c.NoContent(http.StatusNotFound)
		}

		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, rating)
}

func (h Handler) Post(c echo.Context) error {
	req := new(request.Post)
	if err := c.Bind(req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := c.Validate(req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	rating := domain.Rating{
		ProductID: req.ProductID,
		Type:      req.Type,
		Value:     req.Value,
	}

	if err := h.service.Insert(rating); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusCreated)
}

func (h Handler) Put(c echo.Context) error {
	req := new(request.Put)
	if err := c.Bind(req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := c.Validate(req); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	rating := domain.Rating{
		ID:        c.Param("id"),
		ProductID: req.ProductID,
		Type:      req.Type,
		Value:     req.Value,
	}

	if err := h.service.Update(rating); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return c.NoContent(http.StatusNotFound)
		}

		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func (h Handler) Delete(c echo.Context) error {
	if err := h.service.Delete(c.Param("id")); err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}
