package main

import (
	"errors"
	"os"
	"time"

	"github.com/afzalabbasi/QXIP/controller"
	"github.com/sirupsen/logrus"
)

// main method validate env variables, set log level and trigger job function every one second.
func main() {
	hasValidEnvVariables()
	configureLocalFileSystemHook()

	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			logrus.Debugln("Tick at", t)
			controller.QXIPJob()

		}
	}
}

// hasValidEnvVariables method validates environment variables
func hasValidEnvVariables() {

	_, present := os.LookupEnv("URL")
	if !present {
		panic(errors.New("please Provide Valid Loki API URL"))
	}

	_, present = os.LookupEnv("SPEED")
	if !present {
		panic(errors.New("please Provide Valid SPEED PER SECOND"))
	}

	_, present = os.LookupEnv("HEADER")
	if !present {
		panic(errors.New("please Provide Valid Header Information"))
	}

	_, present = os.LookupEnv("LOG_LEVEL")
	if !present {
		panic(errors.New("please Provide Valid LOG_LEVEL"))
	}

}

// configureLocalFileSystemHook method set log level.
func configureLocalFileSystemHook() {
	var level logrus.Level
	loglevel := os.Getenv("LOG_LEVEL")
	if loglevel == "debug" {
		level = logrus.DebugLevel
		logrus.SetLevel(level)
	}
	if loglevel == "error" {
		level = logrus.ErrorLevel
		logrus.SetLevel(level)
	}
	if loglevel == "warn" {
		level = logrus.WarnLevel
		logrus.SetLevel(level)
	}
	if loglevel == "info" {
		level = logrus.InfoLevel
		logrus.SetLevel(level)
	}

}
