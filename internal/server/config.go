package server

type Config struct {
	Host string //`mapstructure:"SERVER_HTTP_CONFIG_HOST"`
	Port int    //`mapstructure:"SERVER_HTTP_CONFIG_PORT"`
}
