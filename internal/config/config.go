package config

import (
	"flag"
	"person-predicator/internal/database"
	"person-predicator/internal/logger"
	"person-predicator/internal/server"

	"github.com/spf13/viper"
)

type Config struct {
	Logger   *logger.Config
	Database *database.Config
	Server   *server.Config
}

func MustLoadConfig() *Config {
	var cfg Config
	path := flag.String("confpath", "./", "path to config file")
	flag.Parse()

	viper.Reset()
	viper.AddConfigPath(*path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	if err := viper.Unmarshal(&cfg.Server); err != nil {
		panic(err.Error())
	}
	if err := viper.Unmarshal(&cfg.Database); err != nil {
		panic(err.Error())
	}
	if err := viper.Unmarshal(&cfg.Logger); err != nil {
		panic(err.Error())
	}

	return &cfg
}
