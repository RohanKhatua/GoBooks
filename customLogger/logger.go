package customLogger

import (
	"log"
	"os"
)

type LogLevel int

const (
	DebugLevel LogLevel = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
	PanicLevel
)

type Logger struct {
	debugLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	fatalLogger   *log.Logger
	panicLogger   *log.Logger
}

func NewLogger() *Logger {
	return &Logger{
		debugLogger:   log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		infoLogger:    log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warningLogger: log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		fatalLogger:   log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile),
		panicLogger:   log.New(os.Stderr, "PANIC: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *Logger) Debug(message string) {
	l.debugLogger.Println(message)
}

func (l *Logger) Info(message string) {
	l.infoLogger.Println(message)
}

func (l *Logger) Warning(message string) {
	l.warningLogger.Println(message)
}

func (l *Logger) Error(message string) {
	l.errorLogger.Println(message)
}

func (l *Logger) Fatal(message string) {
	l.fatalLogger.Fatalln(message)
}

func (l *Logger) Panic(message string) {
	l.panicLogger.Panicln(message)
}