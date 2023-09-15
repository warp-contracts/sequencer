package controller

import (
	"os"

	"github.com/cometbft/cometbft/libs/log"

	"github.com/sirupsen/logrus"
	syncerlogger "github.com/warp-contracts/syncer/src/utils/logger"
)

type LoggerWriter struct {
	level  logrus.Level
	logger log.Logger
}

func (l *LoggerWriter) Write(p []byte) (n int, err error) {
	switch l.level {
	case logrus.DebugLevel:
		l.logger.Debug(string(p[:len(p)-1]))
	case logrus.InfoLevel:
		l.logger.Info(string(p[:len(p)-1]))
	default:
		l.logger.Error(string(p[:len(p)-1]))
	}
	return len(p), nil
}

func InitLogger(logger log.Logger, logLevel string) {
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic(err)
	}

	l := logrus.New()
	l.SetLevel(level)
	l.SetOutput(os.Stdout)

	formatter := &logrus.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: true,
	}
	l.SetFormatter(formatter)

	writer := &LoggerWriter{
		level:  level,
		logger: logger.With("module", "arweave"),
	}
	l.SetOutput(writer)

	syncerlogger.InitWithLogger(l)
}
