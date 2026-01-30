package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func initMongoConfig() {
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.user", "admin")
	viper.SetDefault("database.pass", "admin123")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "27017")
	viper.SetDefault("database.user", "admin")
	viper.SetDefault("database.pass", "admin123")
	viper.SetDefault("database.name", "ratings_db")

	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	user := viper.GetString("database.user")
	pass := viper.GetString("database.pass")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s", user, pass, host, port, user)
	viper.SetDefault("database.uri", uri)

}
