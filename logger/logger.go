package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	logger *logrus.Logger
}

func New() Logger {
	l := Logger {
		logger: logrus.New(),
	}
	
	l.logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	l.logger.SetOutput(os.Stdout)
	l.logger.SetLevel(logrus.TraceLevel)

	return l
}

func (l *Logger) Trace(s string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Trace(s)
}

func (l *Logger) Debug(s string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Debug(s)
}

func (l *Logger) Info(s string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Info(s)
}

func (l *Logger) Warn(s string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Warn(s)
}

func (l *Logger) Error(s string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Error(s)
}

func (l *Logger) Fatal(s string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Fatal(s)
}

func (l *Logger) Panic(s string, fields map[string]interface{}) {
	l.logger.WithFields(fields).Panic(s)
}


