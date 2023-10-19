package config

import (
	"flag"
	"fmt"
	"person-predicator/internal/database"
	"person-predicator/internal/logger"
	"person-predicator/internal/server"

	"github.com/spf13/viper"
)

type Config struct {
	LoggerConfig *logger.Config
	Database     *database.Config
	Server       *server.Config
}

func MustLoadConfig() *Config {
	var cfg Config
	path := flag.String("confpath", "./", "path to config file")
	flag.Parse()

	viper.Reset()
	viper.AddConfigPath(*path)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")


	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err.Error())
	}

	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()
	// viper.SetEnvPrefix("server")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}


	// viper.Reset()
	// viper.AddConfigPath(*path)
	// viper.SetConfigName("app")
	// viper.SetConfigType("env")


	fmt.Println(viper.Get("HOST"))
	fmt.Println(viper.GetInt("PORT"))
	fmt.Println(viper.GetString("DATABASE_CONFIG_HOST"))
	fmt.Println(viper.GetInt("DATABASE_CONFIG_PORT"))
	fmt.Println(viper.GetString("DATABASE_CONFIG_NAME"))

	return &cfg
}
