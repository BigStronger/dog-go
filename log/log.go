package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func New(config *Config) (*zap.Logger, error) {
	logWriter, err := getLogWriter()
	if err != nil {
		return nil, err
	}
	encoder := getEncoder()
	level, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		return nil, err
	}
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level),
		zapcore.NewCore(encoder, zapcore.AddSync(logWriter), level),
	)
	logger := zap.New(core, zap.AddCaller())
	return logger, nil
}

func getLogWriter() (*rotatelogs.RotateLogs, error) {
	currDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	logPath := currDir + "/logs/app.log"
	hook, err := rotatelogs.New(
		logPath+".%Y%m%d%H",
		rotatelogs.WithLinkName(logPath),
		rotatelogs.WithRotationCount(100),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		return nil, err
	}
	return hook, nil
}

func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewJSONEncoder(config)
}
