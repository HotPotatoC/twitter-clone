package logger

import (
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var M *zap.SugaredLogger

func Init(debug bool) {
	var consoleEncoder, jsonEncoder zapcore.EncoderConfig

	var level zapcore.Level

	if debug {
		consoleEncoder = zap.NewDevelopmentEncoderConfig()
		jsonEncoder = zap.NewDevelopmentEncoderConfig()
		level = zap.DebugLevel
	} else {
		consoleEncoder = zap.NewProductionEncoderConfig()
		jsonEncoder = zap.NewProductionEncoderConfig()

		consoleEncoder.EncodeTime = zapcore.ISO8601TimeEncoder
		jsonEncoder.EncodeTime = zapcore.ISO8601TimeEncoder
		level = zapcore.Level(zapcore.WarnLevel)
	}

	consoleEncoder.EncodeLevel = nil

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(consoleEncoder),
		zapcore.AddSync(colorable.NewColorableStdout()),
		level,
	)

	logger := zap.New(core, zap.WithClock(zapcore.DefaultClock))
	defer logger.Sync()

	if debug {
		logger = logger.WithOptions(zap.AddCaller())
	}

	M = logger.Sugar()
}
