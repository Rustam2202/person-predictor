package database

type Config struct {
	Host     string `mapstructure:"DATABASE_HOST"`
	Port     int    `mapstructure:"DATABASE_PORT"`
	User     string `mapstructure:"DATABSE_USER"`
	Password string `mapstructure:"DATABAS_PASSWORD"`
	Name     string `mapstructure:"DATABASE_NAME"`
}
