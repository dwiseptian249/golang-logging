package golanglogging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello Guys")
	logrus.Error("This is Error")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello Guys")
	logrus.Warn("This is Warn")
}
