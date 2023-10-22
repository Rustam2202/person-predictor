package logger

type Config struct {
	Encoding         string   `mapstructure:"ENCODING"`
	Level            string   `mapstructure:"LEVEL"`
	OutputPaths      []string `mapstructure:"OUTPUT_PATHS"`
	ErrorOutputPaths []string `mapstructure:"ERROR_OUTPUT_PATHS"`
}
