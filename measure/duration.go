package measure

import (
	"github.com/sirupsen/logrus"
	"time"
)

func LogDurationFrom(lvl logrus.Level, start time.Time, msg string) {
	logrus.WithField("time", time.Since(start)).
		WithField("type", "benchmark").
		Log(lvl, msg)
}
