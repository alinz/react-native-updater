package logme

import (
	"github.com/Sirupsen/logrus"
)

var logger *logrus.Logger

func init() {
	logger = logrus.New()
	logger.Level = logrus.InfoLevel
	logger.Formatter = &logrus.TextFormatter{ForceColors: true, FullTimestamp: true}
}

//Debug logs debug
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

//Info logs info
func Info(args ...interface{}) {
	logger.Info(args...)
}

//Warn logs warn
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

//Fatal logs fatal
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}
