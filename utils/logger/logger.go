package logger

import (
	"os"

	"github.com/jutionck/interview-bootcamp-apps/config"
	"github.com/jutionck/interview-bootcamp-apps/utils/model"
	"github.com/sirupsen/logrus"
)

type MyLogger interface {
	InitializeLogger() error
	LogInfo(requestLog model.RequestLog)
	LogWarning(requestLog model.RequestLog)
	LogError(requestLog model.RequestLog)
}

type myLogger struct {
	cfg config.FileConfig
	log *logrus.Logger
}

func (m *myLogger) InitializeLogger() error {
	file, err := os.OpenFile(m.cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	m.log = logrus.New()
	m.log.SetOutput(file)
	return nil
}

func (m *myLogger) LogInfo(requestLog model.RequestLog) {
	m.log.Info(requestLog)
}

func (m *myLogger) LogWarning(requestLog model.RequestLog) {
	m.log.Warn(requestLog)
}

func (m *myLogger) LogError(requestLog model.RequestLog) {
	m.log.Error(requestLog)
}

func NewMyLogger(cfg config.FileConfig) MyLogger {
	return &myLogger{cfg: cfg}
}
