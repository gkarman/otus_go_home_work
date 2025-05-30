package logger

import (
	"fmt"
	"strings"
	"time"
)

const (
	LevelError = 10
	LevelWarn  = 20
	LevelInfo  = 30
	LevelDebug = 40
)

var levelMap = map[string]int{
	"error": LevelError,
	"warn":  LevelWarn,
	"info":  LevelInfo,
	"debug": LevelDebug,
}

type Logger struct {
	level int
}

func New(level string) *Logger {
	parsedLevel := parseLevel(level)
	return &Logger{level: parsedLevel}
}

func parseLevel(level string) int {
	if lvl, ok := levelMap[strings.ToLower(level)]; ok {
		return lvl
	}
	return LevelInfo
}

func (l *Logger) log(level int, prefix string, msg string) {
	if l.level < level {
		return
	}
	timestamp := time.Now().Format("2025-01-01 01:01:01")
	fmt.Printf("[%s] [%s] %s\n", timestamp, prefix, msg)
}

func (l *Logger) Info(msg string) {
	l.log(LevelInfo, "INFO", msg)
}

func (l *Logger) Warn(msg string) {
	l.log(LevelWarn, "WARN", msg)
}

func (l *Logger) Error(msg string) {
	l.log(LevelError, "ERROR", msg)
}

func (l *Logger) Debug(msg string) {
	l.log(LevelDebug, "DEBUG", msg)
}
