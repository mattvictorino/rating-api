package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/mattvictorino/rating-api/api/handler"
	"github.com/mattvictorino/rating-api/api/validator"
	"github.com/mattvictorino/rating-api/internal/database"
	"github.com/mattvictorino/rating-api/internal/repository"
	"github.com/mattvictorino/rating-api/internal/service"
	"github.com/spf13/viper"
)

type Server struct {
	handler     handler.Handler
	infoHandler handler.Info
}

func NewServer() Server {
	mongoDB, err := database.NewMongo()
	if err != nil {
		log.Error("Failed to connect to MongoDB:", err)
	}

	ratingRepository := repository.NewRatingRepository(mongoDB)
	ratingService := service.NewRatingService(ratingRepository)
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
