package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "Tio").Info("Hello Tio")

	logger.WithField("username", "Dwi").
		WithField("name", "Dwi Septian").
		Info("Hello Guys")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "Dwi",
		"name":     "Dwi Septian",
	}).Info("Welcome to Paradise")
}
