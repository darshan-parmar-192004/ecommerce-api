package logger

import (
	"backend/internal/constants"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.SugaredLogger

func Init() error {
	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:         constants.LoggerEncoding,
		OutputPaths:      []string{constants.LoggerOutputStdout},
		ErrorOutputPaths: []string{constants.LoggerOutputStderr},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        constants.LoggerKeyTime,
			LevelKey:       constants.LoggerKeyLevel,
			NameKey:        constants.LoggerKeyLogger,
			CallerKey:      constants.LoggerKeyCaller,
			MessageKey:     constants.LoggerKeyMsg,
			StacktraceKey:  constants.LoggerKeyStacktrace,
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	zapLog, err := config.Build()
	if err != nil {
		return err
	}

	Log = zapLog.Sugar()
	return nil
}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}
