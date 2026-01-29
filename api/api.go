package api

import (
	"github.com/labstack/echo/v4"
	"github.com/mattvictorino/rating-api/api/handler"
	"github.com/mattvictorino/rating-api/api/validator"
	"github.com/mattvictorino/rating-api/service"
	"github.com/spf13/viper"
)

type Server struct {
	handler     handler.Handler
	infoHandler handler.Info
}

func NewServer() Server {
	ratingService := service.NewRatingService()
	handler := handler.NewHandler(ratingService)

	return Server{
		handler: handler,
	}
}

func (s Server) Start() {
	echoServer := echo.New()
	echoServer.Validator = validator.NewCustomValidator()

	s.registerRoutes(echoServer)

	echoServer.Logger.Fatal(echoServer.Start(":" + viper.GetString("server.port")))
}
