package proposal

import "github.com/cometbft/cometbft/libs/log"

type LoggerMock struct {
	Msg string
}

func (logger *LoggerMock) Debug(msg string, keyvals ...interface{}) {
	logger.Msg = msg
}

func (logger *LoggerMock) Info(msg string, keyvals ...interface{}) {
	logger.Msg = msg
}

func (logger *LoggerMock) Error(msg string, keyvals ...interface{}) {
	logger.Msg = msg
}

func (logger *LoggerMock) With(keyvals ...interface{}) log.Logger {
	return logger
}
