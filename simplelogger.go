// Package simplelogger only has two methods:
// Debug() and Info()
// These two conver all the essentials needed for a logging package.
package simplelogger

import (
	"io"
	"fmt"
)

// The SimpleLogger struct has booleans for both
// information and debug messages, as well as an 
// io.Writer where messages are logged
type SimpleLogger struct {
	// set to true for information messages
	INFO bool

	// set to true for debug messages
	DEBUG bool

	// The underlying io.Writer for messages
	Writer io.Writer
}

// Info() writes information messages only when Info=true
func (l *SimpleLogger) Info(format string, v ...interface{}) {
	if l.INFO {
		l.Writer.Write([]byte(fmt.Sprintf(format, v...)))
	}
}

// Debug() writes information messages only when Debug=true
func (l *SimpleLogger) Debug(format string, v ...interface{}) {
	if l.DEBUG {
		l.Writer.Write([]byte(fmt.Sprintf(format, v...)))
	}
}

