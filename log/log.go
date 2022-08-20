package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
)

func New(config *Config) (API, error) {
	log := logrus.New()
	log.SetReportCaller(true)
	log.SetFormatter(&customizeLogFormat{})
	switch strings.ToLower(config.Level) {
	case "trace":
		log.SetLevel(logrus.TraceLevel)
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	case "fatal":
		log.SetLevel(logrus.FatalLevel)
	case "panic":
		log.SetLevel(logrus.PanicLevel)
	default:
		log.SetLevel(logrus.ErrorLevel)
	}
	log.SetOutput(os.Stdout)
	if config.Dir == nil || *config.Dir == "" {
		log.SetOutput(os.Stdout)
	} else {
		var logPath string
		if config.Dir != nil {
			logPath = *config.Dir + *config.FileName
		} else {
			currDir, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			logPath = currDir + "/logs/app.log"
		}
		logWriter, err := rotatelogs.New(
			logPath+".%Y_%m_%d_%H_%M",
			rotatelogs.WithLinkName(logPath),
			rotatelogs.WithMaxAge(time.Hour*24*90),
			rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
		)
		if err != nil {
			panic(err)
		}
		log.SetOutput(io.MultiWriter(os.Stdout, logWriter))
	}
	return log, nil
}
