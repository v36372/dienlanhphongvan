package logger

import (
	"github.com/Sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	errLogFile = "log/error.log"
)

type Logger struct {
	logger *logrus.Logger
}

func ErrJSONFile() *Logger {
	return newJSONLogger(errLogFile)
}

func newJSONLogger(file string) *Logger {
	return &Logger{
		logger: newJSONLogrus(file),
	}
}

func newJSONLogrus(file string) *logrus.Logger {
	format := logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999999999",
	}
	return &logrus.Logger{
		Out:       fileWriter(file),
		Formatter: &format,
		Level:     logrus.DebugLevel,
	}
}

func fileWriter(log string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   log,
		MaxSize:    100, // megabytes
		MaxBackups: 5,
		MaxAge:     30, // days
	}
}

func (l *Logger) Error(msg string, fields map[string]interface{}) {
	if l.logger == nil {
		return
	}

	l.logger.WithFields(l.convert(fields)).Error(msg)
}

func (l *Logger) Info(msg string, fields map[string]interface{}) {
	if l.logger == nil {
		return
	}
	l.logger.WithFields(l.convert(fields)).Info(msg)
}

func (l *Logger) convert(values map[string]interface{}) logrus.Fields {
	return logrus.Fields(values)
}
