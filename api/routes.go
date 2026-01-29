package api

import "github.com/labstack/echo/v4"

func (s Server) registerRoutes(e *echo.Echo) {
	ratingGroup := e.Group("rating")
	ratingGroup.GET(":id", s.handler.Get)
	ratingGroup.POST("", s.handler.Post)
	ratingGroup.PUT(":id", s.handler.Put)
	ratingGroup.DELETE(":id", s.handler.Delete)

	e.GET("info", s.infoHandler.Info)
}
