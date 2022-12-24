package infrastructure

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// / Logger structure
type Logger struct {
	Zap *zap.SugaredLogger
}

func NewLogger(env Env) Logger {
	config := zap.NewDevelopmentConfig()
	if env.ENVIRONMENT == "local" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	logger, _ := config.Build(zap.Hooks(func(e zapcore.Entry) error {
		return nil
	}))
	sugar := logger.Sugar()

	return Logger{
		Zap: sugar,
	}

}
