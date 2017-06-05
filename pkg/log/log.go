package log

import (
	"os"
	"github.com/sirupsen/logrus"
	"github.com/ridewindx/loong/cmd"
)

func initLog() error {
	level, err := logrus.ParseLevel(cmd.LogLevel)
	if err != nil {
		return err
	}
	logrus.SetLevel(level)

	logrus.SetFormatter(&logrus.TextFormatter{})

	logrus.SetOutput(os.Stdout)

	logrus.AddHook(ContextHook{})

	return nil
}
