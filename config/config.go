package config

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func InitConfig() {
	_ = godotenv.Load()

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetDefault("server.port", 8888)

	initMongoConfig()
}
