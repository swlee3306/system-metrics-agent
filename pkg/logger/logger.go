// Package logger provides standard log output without persistence or retention.
package logger

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

const (
	DebugLevel = "DEBUG"
	InfoLevel  = "INFO"
	WarnLevel  = "WARN"
	ErrorLevel = "ERROR"
	FatalLevel = "FATAL"
)

func Debug(target, message string, args ...any) { write(DebugLevel, target, message, args...) }
func Info(target, message string, args ...any)  { write(InfoLevel, target, message, args...) }
func Warn(target, message string, args ...any)  { write(WarnLevel, target, message, args...) }
func Error(target, message string, args ...any) { write(ErrorLevel, target, message, args...) }

// Fatal records a level; it does not terminate the process.
func Fatal(target, message string, args ...any) { write(FatalLevel, target, message, args...) }

func write(level, target, message string, args ...any) {
	detail := "unknown"
	if pc, file, line, ok := runtime.Caller(2); ok {
		name := "unknown"
		if fn := runtime.FuncForPC(pc); fn != nil {
			name = fn.Name()
		}
		detail = fmt.Sprintf("%s:%s:%d", name, filepath.Base(file), line)
	}
	log.Printf("LogLevel:(%s), Target:(%s), detail:(%s), LogMessage:(%s)",
		level, target, detail, fmt.Sprintf(message, args...))
}
