package glogger

import (
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

// keep files open and write on log file in an interval if it is production level
// delete, openFiles if it is closed.
// close open files when shutting down the file

var openFiles sync.Map

type FileLogger struct {
	level  LogLevel
	name   string
	logger *logrus.Logger
}

// NewFileLogger creates a new file logger
func NewFileLogger(name string) *FileLogger {
	var file *os.File
	existedFile, ok := openFiles.Load(name)
	if !ok {
		newFile, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logrus.Error("Failed to open log file:", err)
		}

		file = newFile
		openFiles.Store(name, newFile)

	} else {
		file = existedFile.(*os.File)
	}

	logger := logrus.New()
	logger.SetOutput(file)

	return &FileLogger{name: name, logger: logger}
}

func (l *FileLogger) Error(args ...any) {
	//if !std.allowed() {
	//	return
	//}
	if !isAllowed(ErrorLevel) {
		return
	}

	l.logger.Error(args...)
}

func (l *FileLogger) Errorf(format string, args ...any) {
	if !isAllowed(ErrorLevel) {
		return
	}

	l.logger.Errorf(format, args...)
}

func (l *FileLogger) Warn(args ...any) {
	if !isAllowed(WarnLevel) {
		return
	}
}

func (l *FileLogger) Warnf(format string, args ...any) {
	if !isAllowed(WarnLevel) {
		return
	}
}

func (l *FileLogger) Info(args ...any) {
	if !isAllowed(InfoLevel) {
		return
	}
}

func (l *FileLogger) Infof(format string, args ...any) {
	if !isAllowed(InfoLevel) {
		return
	}
}

func (l *FileLogger) Debug(args ...any) {
	l.logger.Debug(args...)
}

func (l *FileLogger) Debugf(format string, args ...any) {
	l.logger.Debugf(format, args...)
}

func (l *FileLogger) WithLevel(level LogLevel) *FileLogger {
	l.level = level
	return l
}

func Shutdown() error {
	openFiles.Range(func(key, value any) bool {
		file := value.(*os.File)

		file.Close()
		openFiles.Delete(key)

		return true
	})

	return nil
}
