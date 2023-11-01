package glogger

import (
	"github.com/sirupsen/logrus"
	"log"
)

var std *Logger

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

// prodValidLevel is those levels that are active firing in production
var prodValidLevel = map[LogLevel]bool{
	DebugLevel: false,
	InfoLevel:  false,
	WarnLevel:  true,
	ErrorLevel: true,
	FatalLevel: true,
}

/*
 1. test Logger behaviour on concurrent situation
 2. what happens to custom log file in container environment
 3. Can we implement a better solution for concurrent file logging?
 4. Add sources like network sources
 5. How can I add sentry for errors: define a subscriber mechanism that all clients can subscribe to it
    and the Logger give the prompt of logging for all of them (standard, file, third-party)
    Sentry for instance is a dedicated package that can listens to a specific channel to send received errors
*/

func init() {
	std = New()
}

type Logger struct {
	level       LogLevel
	production  bool
	subscribers SubscriptionList
}

func New() *Logger {
	if std != nil {
		return std
	}

	format := CustomFormatter{}
	logrus.SetFormatter(format)

	return &Logger{
		level:       InfoLevel,
		production:  true,
		subscribers: nil,
	}
}

type CustomFormatter struct {
	prefix string
}

func (f CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	entry.Message = f.prefix + entry.Message
	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05 -07:00",
	}
	return formatter.Format(entry)
}

type Config struct {
	Production bool
}

func SetConfig(config Config) {
	if !config.Production {
		std.production = false
		logrus.SetLevel(logrus.DebugLevel)
	}
}

func isAllowed(level LogLevel) bool {
	if allowed := prodValidLevel[level]; std.production && !allowed {
		return false
	}
	return true
}

func Print(args ...any) {
	log.Println(args...)
}

func Println(format string, args ...any) {
	log.Printf(format, args...)
}

func Log(args ...any) {
	logrus.Println(args...)
}

func Logf(format string, args ...any) {
	logrus.Printf(format, args...)
}

func Info(args ...any) {
	// control the production flag
	if !isAllowed(InfoLevel) {
		return
	}

	logrus.Infoln(args...)
}

func Infof(format string, args ...any) {
	// print the log with format
	if !isAllowed(InfoLevel) {
		return
	}

	logrus.Infof(format, args...)
}

func Debug(args ...any) {
	logrus.Debug(args...)
}

func Debugf(format string, args ...any) {
	logrus.Debugf(format, args...)
}

func Warn(args ...any) {
	// control the production flag
	if !isAllowed(WarnLevel) {
		return
	}

	logrus.Warn(args...)
	std.subscribers.publish(WarnLevel, []byte("error from Logger"))
}

func Warnf(format string, args ...any) {
	// print the log with format
	if !isAllowed(WarnLevel) {
		return
	}

	logrus.Warnf(format, args...)
}

func Error(args ...any) {
	// control the production flag
	if !isAllowed(ErrorLevel) {
		return
	}

	logrus.Error(args...)
	std.subscribers.publish(ErrorLevel, []byte("error from Logger"))
}

func Errorf(format string, args ...any) {
	// print the log with format
	if !isAllowed(ErrorLevel) {
		return
	}

	logrus.Errorf(format, args...)
}

func Fatal(args ...any) {
	// control the production flag
	if !isAllowed(FatalLevel) {
		return
	}

	std.subscribers.publish(FatalLevel, []byte("error from Logger"))
	logrus.Fatal(args...)
}

func Fatalf(format string, args ...any) {
	// print the log with format
	if !isAllowed(FatalLevel) {
		return
	}

	logrus.Fatalf(format, args...)
}
