package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"
)

var stdLogger *log.Logger

// syncWriter wraps io.Writer and calls Sync() after each Write
// to ensure data is flushed to disk immediately.
type syncWriter struct {
	w io.Writer
}

func (sw *syncWriter) Write(p []byte) (n int, err error) {
	n, err = sw.w.Write(p)
	if syncer, ok := sw.w.(interface{ Sync() error }); ok {
		syncer.Sync()
	}
	return n, err
}

func SetupLogger() {
	os.MkdirAll("logs", 0755)

	fileWriter := &lumberjack.Logger{
		Filename:   fmt.Sprintf("./logs/%s.log", time.Now().Format("2006-01-02")),
		MaxSize:    100,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   true,
	}

	stdLogger = log.New(&syncWriter{w: fileWriter}, "", 0)
}

func formatMessage(level, msg string, err error) {
	pc, file, line, ok := runtime.Caller(3)
	funcName := "unknown"
	if ok {
		if fn := runtime.FuncForPC(pc); fn != nil {
			funcName = filepath.Base(fn.Name())
		}
	}

	if err != nil {
		msg = fmt.Sprintf("%s: %v", msg, err)
	}

	ts := time.Now().Format(time.RFC3339Nano)
	stdLogger.Printf("%s | %s | %s:%d | %s | %s",
		ts, level, filepath.Base(file), line, funcName, msg)
}

func Info(msg string, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	formatMessage("INFO", msg, nil)
}

func Warn(msg string, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	formatMessage("WARN", msg, nil)
}

func Error(msg string, err error, v ...interface{}) {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	formatMessage("ERROR", msg, err)
}

// LoggerImpl mengimplementasikan Logger interface dari application layer
type LoggerImpl struct{}

func (l *LoggerImpl) Info(msg string, args ...interface{}) {
	Info(msg, args...)
}

func (l *LoggerImpl) Error(msg string, err error, args ...interface{}) {
	Error(msg, err, args...)
}

func (l *LoggerImpl) Warn(msg string, args ...interface{}) {
	Warn(msg, args...)
}
