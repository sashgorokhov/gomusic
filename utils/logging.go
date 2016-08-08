package utils

import (
	"github.com/Sirupsen/logrus"
	"github.com/sashgorokhov/govk"
	"github.com/x-cray/logrus-prefixed-formatter"
	"io"
	"io/ioutil"
	"log"
	"os"
)

var GomusicLogger = (&logrus.Logger{
	Level:     logrus.DebugLevel,
	Formatter: new(prefixed.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Out:       ioutil.Discard,
}).WithField("prefix", "gomusic")

func SetupLogger(out io.Writer, logger *logrus.Logger) {
	logger.Out = out
	logger.Level = logrus.DebugLevel
}

func SetupLoggers(out io.Writer, loggers []*logrus.Logger) {
	for _, logger := range loggers {
		SetupLogger(out, logger)
	}
}

func GetAllLoggers() []*logrus.Logger {
	return []*logrus.Logger{govk.ApiLogger.Logger, govk.AuthLogger.Logger, GomusicLogger.Logger}
}

func SetupLogging() func() {
	all_loggers := GetAllLoggers()

	file, err := os.OpenFile(LOG_FILENAME, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)

	if err != nil {
		log.Println("Cant setup logging file: " + err.Error())
		SetupLoggers(os.Stderr, all_loggers)
		return nil
	}

	SetupLoggers(file, all_loggers)

	return func() {
		file.Close()
	}
}
