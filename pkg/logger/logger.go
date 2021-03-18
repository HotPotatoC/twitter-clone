package logger

import (
	"github.com/mattn/go-colorable"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger Constructs a new logger
func NewLogger(debug bool) *zap.SugaredLogger {
	atom := zap.NewAtomicLevel()

	if debug {
		atom.SetLevel(zap.DebugLevel)
	} else {
		atom.SetLevel(zap.InfoLevel)
	}

	encoder := zap.NewDevelopmentEncoderConfig()

	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder),
		zapcore.AddSync(colorable.NewColorableStdout()),
		atom,
	))
	defer logger.Sync()

	if debug {
		logger = logger.WithOptions(zap.AddCaller())
		return logger.Sugar()
	}

	sugar := logger.Sugar()

	return sugar
}
