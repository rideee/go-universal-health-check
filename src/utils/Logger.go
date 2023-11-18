package utils

import "fmt"

type Logger struct {
	DebugMode bool
}

// Constructor.
func NewLogger(debugMode bool) *Logger {
	l := new(Logger)
	l.DebugMode = debugMode
	return l
}

// Same as fmt.Println, but also writes the message to the log file.
// TODO: store logs in a file...
func (l *Logger) Println(msg ...any) {
	fmt.Println(msg...)
}

// Same as fmt.Printf, but also writes the message to the log file.
// TODO: store logs in a file...
func (l *Logger) Printf(format string, v ...any) {
	fmt.Printf(format, v...)
}

// Same as fmt.Println, but also writes the message to the log file.
// TODO: store logs in a file...
func (l *Logger) Debug(msg ...any) {
	if !l.DebugMode {
		return
	}
	fmt.Print("DEBUG: ")
	fmt.Println(msg...)
}

// Same as fmt.Printf, but also writes the message to the log file.
// TODO: store logs in a file...
func (l *Logger) Debugf(format string, v ...any) {
	if !l.DebugMode {
		return
	}
	fmt.Print("DEBUG: ")
	fmt.Printf(format, v...)
}

// Same as fmt.Println, but also writes the message to the log file.
// TODO: store logs in a file...
func (l *Logger) Warning(msg ...any) {
	fmt.Print("WARNING: ")
	fmt.Println(msg...)
}

// Same as fmt.Printf, but also writes the message to the log file.
// TODO: store logs in a file...
func (l *Logger) Warningf(format string, v ...any) {
	fmt.Print("WARNING: ")
	fmt.Printf(format, v...)
}

// Same as fmt.Println, but also writes the message to the log file.
// TODO: store logs in a file...
func (l *Logger) Error(msg ...any) {
	fmt.Print("ERROR: ")
	fmt.Println(msg...)
}

// Same as fmt.Printf, but also writes the message to the log file.
// TODO: store logs in a file...
func (l *Logger) Errorf(format string, v ...any) {
	fmt.Print("ERROR: ")
	fmt.Printf(format, v...)
}
