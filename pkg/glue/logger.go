package glue

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func DefaultFileLogger(cfg *Log) *Logger {
	log := logrus.New()
	log.SetFormatter(GetFormatter(cfg.LogTimestampFormat))
	setLogLvl(log, cfg.LogLevel)
	setLogOutput(log, cfg.LogFile)

	var logger = &Logger{log}
	return logger
}

func GetFormatter(timeForm string) logrus.Formatter {
	jf := new(logrus.JSONFormatter)
	jf.TimestampFormat = timeForm
	return jf
}

func setLogOutput(log *logrus.Logger, file string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		logrus.Panic(err.Error())
	}
	log.SetOutput(io.MultiWriter(os.Stdout, f))
}

func setLogLvl(log *logrus.Logger, logLvl string) {
	lvl, err := logrus.ParseLevel(logLvl)
	if err != nil {
		logrus.Panic(err.Error())
	}
	log.SetLevel(lvl)
}
