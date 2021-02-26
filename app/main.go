package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {

	logger := Logger()

	for {

		switch i := rand.Intn(4); i {
		case 0:
			logger.WithFields(logrus.Fields{"file": "main.go"}).Debug("this is a DEBUG message")
		case 1:
			logger.WithFields(logrus.Fields{"file": "main.go"}).Error("this is an ERROR message")
		case 2:
			logger.WithFields(logrus.Fields{"file": "main.go"}).Info("this is an INFO message")
		case 3:
			logger.WithFields(logrus.Fields{"file": "main.go"}).Warn("this is an WARN message")
		default:
			logger.WithFields(logrus.Fields{"file": "main.go"}).Debug("this is a DEBUG message")
		}

		time.Sleep(1 * time.Second)
	}
}

func Logger() *logrus.Logger {
	logger := logrus.New()
	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.DebugLevel)
	log.SetOutput(logger.Writer())
	logger.SetOutput(io.MultiWriter(os.Stdout))
	return logger
}
