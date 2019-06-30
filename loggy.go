// Package loggy provides a simple logging library for Go.
package loggy // import "go.gophers.dev/pkgs/loggy"

import (
	"log"
	"os"
)

// A Logger is used to write log lines.
//
// There are four methods for outputting a log message, each being associated
// with a different degree of urgency.
type Logger interface {
	Tracef(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
}

type logger struct {
	tag    string
	output Printer
}

// A Printer is anything that can be used to Printf log messages.
//
// The default underlying implementation of Logger provided by the loggy
// package defers to the standard library log package, which in turn provides
// an implementation of Printer that writes to an output file.
type Printer interface {
	Printf(string, ...interface{})
}

// Open creates a standard library *log.Log with the given file as the
// underlying output device.
func Open(file *os.File) Printer {
	return log.New(file, "", log.LstdFlags)
}

func tag(prefix string) string {
	return "[" + prefix + "]"
}

// Options allow for configuring a Logger, with the given Prefix and Output
// mechanism.
type Options struct {
	Prefix string
	Output Printer
}

// New creates a Logger which sends output to os.Stderr with each log line
// tagged with the given prefix.
//
// To control the output destination, use Configure to create the Logger instead.
func New(prefix string) Logger {
	return Configure(Options{
		Prefix: prefix,
		Output: Open(os.Stderr),
	})
}

// Configure creates a Logger with the configured Options.
//
// This allows for controlling the output destination of the created Logger.
func Configure(opts Options) Logger {
	return &logger{
		tag:    tag(opts.Prefix),
		output: opts.Output,
	}
}

// Discard creates a Logger which throws away each statement.
func Discard() Logger {
	devNull := os.NewFile(0, os.DevNull)
	return Configure(Options{
		Output: Open(devNull),
	})
}

func (l logger) Tracef(format string, a ...interface{}) {
	l.printf("TRACE", format, a...)
}

func (l logger) Infof(format string, a ...interface{}) {
	l.printf("INFO ", format, a...)
}

func (l logger) Warnf(format string, a ...interface{}) {
	l.printf("WARN ", format, a...)
}

func (l logger) Errorf(format string, a ...interface{}) {
	l.printf("ERROR", format, a...)
}

func (l logger) printf(level, format string, a ...interface{}) {
	prefixedFmt := level + " " + l.tag + " " + format
	l.output.Printf(prefixedFmt, a...)
}
