package util

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Log() *logrus.Logger {
	var log = logrus.New()
	log.SetLevel(logrus.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	return log
}
