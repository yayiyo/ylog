package ylog

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

type Level uint8

const (
	// debug level logs are typically voluminous, and are usually disabled in production.
	LevelDebug Level = iota
	// info level is the default logging priority.
	LevelInfo
	// warning logs are more important than info, but don't need individual human review.
	LevelWarning
	// error logs are higu-priority. if an application is running smoothly,
	// it shouldn't be generated ant error-level logs.
	LevelError
	// panic logs a message, then panics.
	LevelPanic
	// fatal logs a message, then calls os.Exit(1).
	LevelFatal
)

var LevelNameMapping = map[Level]string{
	LevelDebug:   "DEBUG",
	LevelInfo:    "INFO",
	LevelWarning: "WARNING",
	LevelError:   "ERROR",
	LevelPanic:   "PANIC",
	LevelFatal:   "FATAL",
}

var errUnmarshalNilLevel = errors.New("can't unmarshal nil *Level")

func (l *Level) unmarshalText(text []byte) bool {
	switch string(text) {
	case "debug", "DEBUG":
		*l = LevelDebug
	case "info", "INFO", "": // make the zero value useful
		*l = LevelInfo
	case "warn", "WARN":
		*l = LevelWarning
	case "error", "ERROR":
		*l = LevelError
	case "panic", "PANIC":
		*l = LevelPanic
	case "fatal", "FATAL":
		*l = LevelFatal
	default:
		return false
	}
	return true
}

func (l *Level) UnmarshalText(text []byte) error {
	if l == nil {
		return errUnmarshalNilLevel
	}

	if !l.unmarshalText(text) && !l.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized level: %q", text)
	}
	return nil
}

type options struct {
	output        io.Writer
	level         Level
	stdLevel      Level
	formatter     Formatter
	disableCaller bool
}

type Option func(*options)

func initOptions(opts ...Option) (o *options) {
	o = &options{}
	for _, opt := range opts {
		opt(o)
	}
	if o.output == nil {
		o.output = os.Stderr
	}

	if o.formatter == nil {
		o.formatter = &TextFormatter{}
	}
	return
}

func WithOutput(output io.Writer) Option {
	return func(o *options) {
		o.output = output
	}
}

func WithLevel(level Level) Option {
	return func(o *options) {
		o.level = level
	}
}

func WithStdLevel(level Level) Option {
	return func(o *options) {
		o.stdLevel = level
	}
}

func WithFormatter(formatter Formatter) Option {
	return func(o *options) {
		o.formatter = formatter
	}
}

func WithDisableCaller(caller bool) Option {
	return func(o *options) {
		o.disableCaller = caller
	}
}
