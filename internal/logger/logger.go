package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func MustConfigLogger(cfg Config) {
	congig := zap.NewProductionConfig()
	congig.Encoding = cfg.Encoding
	err := congig.Level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		panic(err.Error())
	}
	congig.OutputPaths = cfg.OutputPaths
	congig.ErrorOutputPaths = cfg.ErrorOutputPaths
	congig.EncoderConfig = zap.NewProductionEncoderConfig()
	Logger, err = congig.Build()
	if err != nil {
		panic(err.Error())
	}
}
