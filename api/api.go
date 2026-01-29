package api

import (
	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (s Server) Start() {
	echoServer := echo.New()
	echoServer.Logger.Fatal(echoServer.Start(":" + viper.GetString("server.port")))
}
