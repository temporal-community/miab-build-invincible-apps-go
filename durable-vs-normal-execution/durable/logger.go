package durable

import (
	"fmt"
	"strings"

	"go.temporal.io/sdk/workflow"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warn
	Error
)

type SimpleLogger struct {
	level LogLevel
	ctx   workflow.Context
}

func NewSimpleLogger(level LogLevel) *SimpleLogger {
	return &SimpleLogger{
		level: level,
	}
}

// Debug logs debug level messages
func (l *SimpleLogger) Debug(msg string, keyvals ...interface{}) {
	if l.level <= Debug && !isReplaying(l.ctx) {
		l.log(msg)
	}
}

// Info logs info level messages
func (l *SimpleLogger) Info(msg string, keyvals ...interface{}) {
	if l.level <= Info && !isReplaying(l.ctx) {
		l.log(msg)
	}
}

// Warn logs warning level messages
func (l *SimpleLogger) Warn(msg string, keyvals ...interface{}) {
	if l.level <= Warn && !isReplaying(l.ctx) {
		l.log(msg)
	}
}

// Error logs error level messages
func (l *SimpleLogger) Error(msg string, keyvals ...interface{}) {
	if l.level <= Error && !isReplaying(l.ctx) {
		l.log(msg)
	}
}

// With returns a logger with additional key-value pairs and context
func (l *SimpleLogger) With(keyvals ...interface{}) *SimpleLogger {
	// Check if context is being passed
	for i := 0; i < len(keyvals); i += 2 {
		if ctx, ok := keyvals[i].(workflow.Context); ok {
			return &SimpleLogger{
				level: l.level,
				ctx:   ctx,
			}
		}
	}
	return l
}

func isReplaying(ctx workflow.Context) bool {
	if ctx == nil {
		return false
	}
	return workflow.IsReplaying(ctx)
}

func (l *SimpleLogger) log(msg string, keyvals ...interface{}) {
	if len(keyvals) == 0 {
		fmt.Println(msg)
		return
	}

	// Handle format strings
	if len(keyvals) == 1 {
		fmt.Printf(msg, keyvals...)
		fmt.Println()
		return
	}

	// Handle key-value pairs
	var pairs []string
	for i := 0; i < len(keyvals); i += 2 {
		key := fmt.Sprintf("%v", keyvals[i])
		var value string
		if i+1 < len(keyvals) {
			value = fmt.Sprintf("%v", keyvals[i+1])
		}
		pairs = append(pairs, fmt.Sprintf("%s=%s", key, value))
	}

	if len(pairs) > 0 {
		fmt.Printf("%s %s\n", msg, strings.Join(pairs, " "))
	} else {
		fmt.Println(msg)
	}
}
