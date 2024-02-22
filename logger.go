package ylog

import (
	"fmt"
	"io"
	"os"
	"sync"
	"unsafe"
)

var std = New()

type logger struct {
	opt       *options
	mu        sync.Mutex
	entryPool *sync.Pool
}

func New(opts ...Option) *logger {
	log := &logger{opt: initOptions(opts...)}
	log.entryPool = &sync.Pool{New: func() any { return entry(log) }}
	return log
}

func StdLogger() *logger {
	return std
}

func SetOptions(opts ...Option) {
	std.SetOptions(opts...)
}

func (l *logger) SetOptions(opts ...Option) {
	l.mu.Lock()
	defer l.mu.Unlock()

	for _, opt := range opts {
		opt(l.opt)
	}
}

func Writer() io.Writer {
	return std
}

func (l *logger) Writer() io.Writer {
	return l
}

func (l *logger) Write(data []byte) (int, error) {
	l.entry().write(l.opt.stdLevel, "", *(*string)(unsafe.Pointer(&data)))
	return 0, nil
}

func (l *logger) entry() *Entry {
	return l.entryPool.Get().(*Entry)
}

func (l *logger) Debug(args ...any) {
	l.entry().write(LevelDebug, "", args...)
}

func (l *logger) Debugf(format string, args ...any) {
	l.entry().write(LevelDebug, format, args...)
}

func (l *logger) Info(args ...any) {
	l.entry().write(LevelInfo, "", args...)
}

func (l *logger) Infof(format string, args ...any) {
	l.entry().write(LevelInfo, format, args...)
}

func (l *logger) Warn(args ...any) {
	l.entry().write(LevelWarning, "", args...)
}

func (l *logger) Warnf(format string, args ...any) {
	l.entry().write(LevelWarning, format, args...)
}

func (l *logger) Error(args ...any) {
	l.entry().write(LevelError, "", args...)
}

func (l *logger) Errorf(format string, args ...any) {
	l.entry().write(LevelError, format, args...)
}

func (l *logger) Panic(args ...any) {
	l.entry().write(LevelPanic, "", args...)
	panic(fmt.Sprint(args...))
}

func (l *logger) Panicf(format string, args ...any) {
	l.entry().write(LevelPanic, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func (l *logger) Fatal(args ...any) {
	l.entry().write(LevelFatal, "", args...)
	os.Exit(1)
}

func (l *logger) Fatalf(format string, args ...any) {
	l.entry().write(LevelFatal, format, args...)
	os.Exit(1)
}

// std logger
func Debug(args ...any) {
	std.entry().write(LevelDebug, "", args...)
}

func Debugf(format string, args ...any) {
	std.entry().write(LevelDebug, format, args...)
}

func Info(args ...any) {
	std.entry().write(LevelInfo, "", args...)
}

func Infof(format string, args ...any) {
	std.entry().write(LevelInfo, format, args...)
}

func Warn(args ...any) {
	std.entry().write(LevelWarning, "", args...)
}

func Warnf(format string, args ...any) {
	std.entry().write(LevelWarning, format, args...)
}

func Error(args ...any) {
	std.entry().write(LevelError, "", args...)
}

func Errorf(format string, args ...any) {
	std.entry().write(LevelError, format, args...)
}

func Panic(args ...any) {
	std.entry().write(LevelPanic, "", args...)
	panic(fmt.Sprint(args...))
}

func Panicf(format string, args ...any) {
	std.entry().write(LevelPanic, format, args...)
	panic(fmt.Sprintf(format, args...))
}

func Fatal(args ...any) {
	std.entry().write(LevelFatal, "", args...)
	os.Exit(1)
}

func Fatalf(format string, args ...any) {
	std.entry().write(LevelFatal, format, args...)
	os.Exit(1)
}
