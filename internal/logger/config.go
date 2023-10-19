package logger

type Config struct {
	Encoding         string   `yaml:"encoding"`
	Level            string   `yaml:"level"`
	OutputPaths      []string `yaml:"outputPaths"`
	ErrorOutputPaths []string `yaml:"errorOutputPaths"`
}
