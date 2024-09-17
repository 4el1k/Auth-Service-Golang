package main

import (
	"Avito-Test-Backend-Golang/internal/app"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {

	log := logrus.New()
	log.SetLevel(logrus.DebugLevel)
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.DateTime,
		FullTimestamp:   true,
	})

	app := app.NewApp(log)

	if err := app.Run(); err != nil {
		os.Exit(1)
	}
}
