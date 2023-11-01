package glogger

type ISubscriber interface {
	GetClosure() func(message []byte) error
}

type ILogger interface {
	Error(args ...any)
	Errorf(format string, args ...any)
	Warn(args ...any)
	Warnf(format string, args ...any)
	Info(args ...any)
	Infof(format string, args ...any)
	Debug(args ...any)
	Debugf(format string, args ...any)
}
