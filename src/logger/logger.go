package logger

import (
	"catalog/src/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(conf *config.LoggerConfig) (*zap.Logger, error) {
	level := zap.NewAtomicLevel()

	err := level.UnmarshalText([]byte(conf.LogLevel))
	if err != nil {
		return nil, err
	}

	encoding := "json"
	encodeLevel := zapcore.LowercaseLevelEncoder
	if conf.DevMode {
		level.SetLevel(zapcore.DebugLevel)
		encoding = "console"
		encodeLevel = zapcore.LowercaseColorLevelEncoder
	}

	logconf := zap.Config{
		Level:       level,
		Encoding:    encoding,
		Development: conf.DevMode,
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "severity",
			TimeKey:        "timestamp",
			EncodeLevel:    encodeLevel,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
		},
	}
	log, err := logconf.Build()
	if err != nil {
		return nil, err
	}

	return log, nil
}
