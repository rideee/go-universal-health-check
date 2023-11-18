package utils

import "fmt"

// Same as fmt.Println, but also writes the message to the log file.
// TODO: store logs in a file...
func Println(msg ...any) {
	fmt.Println(msg...)
}

// Same as fmt.Printf, but also writes the message to the log file.
// TODO: store logs in a file...
func Printf(format string, v ...any) {
	fmt.Printf(format, v...)
}

// Same as fmt.Println, but also writes the message to the log file.
// TODO: store logs in a file...
func Debug(msg ...any) {
	fmt.Println(msg...)
}

// Same as fmt.Printf, but also writes the message to the log file.
// TODO: store logs in a file...
func Debugf(format string, v ...any) {
	fmt.Printf(format, v...)
}

// Same as fmt.Println, but also writes the message to the log file.
// TODO: store logs in a file...
func Warning(msg ...any) {
	fmt.Println(msg...)
}

// Same as fmt.Printf, but also writes the message to the log file.
// TODO: store logs in a file...
func Warningf(format string, v ...any) {
	fmt.Printf(format, v...)
}

// Same as fmt.Println, but also writes the message to the log file.
// TODO: store logs in a file...
func Error(msg ...any) {
	fmt.Println(msg...)
}

// Same as fmt.Printf, but also writes the message to the log file.
// TODO: store logs in a file...
func Errorf(format string, v ...any) {
	fmt.Printf(format, v...)
}
