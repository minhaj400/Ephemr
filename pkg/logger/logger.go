package logger

import (
	"log"
	"os"
)

type logger struct {
	debug *log.Logger
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
}

func NewLogger() *logger {
	return &logger{
		debug: log.New(os.Stdout, "DEBUG: ", log.LstdFlags),
		info:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
		warn:  log.New(os.Stdout, "WARN: ", log.LstdFlags),
		error: log.New(os.Stderr, "ERROR: ", log.LstdFlags),
	}
}

func (l *logger) Debug(v ...any) {
	l.debug.Println(v...)
}

func (l *logger) Info(v ...any) {
	l.info.Println(v...)
}

func (l *logger) Warn(v ...any) {
	l.warn.Println(v...)
}

func (l *logger) Error(v ...any) {
	l.error.Println(v...)
}
