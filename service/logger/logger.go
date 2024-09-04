package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var (
	rootLogger *zap.SugaredLogger
)

func RootLogger() (*zap.SugaredLogger, error) {
	if rootLogger != nil {
		return rootLogger, nil
	}

	var logger *zap.Logger
	var err error
	if logger, err = zap.NewDevelopment(); err != nil {
		return nil, err
	}
	rootLogger = logger.Sugar()
	return rootLogger, nil
}

func NamedLogger(name string) (*zap.SugaredLogger, error) {
	root, err := RootLogger()
	if err != nil {
		return nil, err
	}
	return root.Named(name), nil
}

type LogWrapper interface {
	Logger() *zap.SugaredLogger
	SetLogger(logger *zap.SugaredLogger)
	LoggerName() string

	Infow(msg string, args ...interface{})
	Errorw(msg string, args ...interface{})
	Debugw(msg string, args ...interface{})
}

type logWrapper struct {
	logger *zap.SugaredLogger
	name   string
}

func NewLogWrapper(name string) *logWrapper {
	return &logWrapper{
		logger: nil,
		name:   name,
	}
}

func (lw *logWrapper) Logger() *zap.SugaredLogger {
	return lw.logger
}

func (lw *logWrapper) SetLogger(logger *zap.SugaredLogger) {
	lw.logger = lw.setLoggerName(logger)
}

func (lw *logWrapper) setLoggerName(logger *zap.SugaredLogger) *zap.SugaredLogger {
	if logger == nil {
		fmt.Println("ERROR: tried to set name '", lw.name, "' on nil logger!")
		return logger
	}
	return logger.Named(lw.name)
}

func (lw *logWrapper) LoggerName() string {
	return lw.name
}

func (lw *logWrapper) Infow(msg string, args ...interface{}) {
	if lw.logger != nil {
		lw.logger.Infow(msg, args...)
	}
}

func (lw *logWrapper) Errorw(msg string, args ...interface{}) {
	if lw.logger != nil {
		lw.logger.Errorw(msg, args...)
	}
}

func (lw *logWrapper) Debugw(msg string, args ...interface{}) {
	if lw.logger != nil {
		lw.logger.Debugw(msg, args...)
	}
}
