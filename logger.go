package obeliskagi

import (
	"log"
	"os"
)

// StandardLogger represents a contract for general loggers.
type StandardLogger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}

type defaultLogger struct {
	enabled     bool
	InfoLogger  *log.Logger
	FatalLogger *log.Logger
	PanicLogger *log.Logger
}

func newLogger(enabled bool) *defaultLogger {
	return &defaultLogger{
		enabled:     enabled,
		InfoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		FatalLogger: log.New(os.Stderr, "FATAL: ", log.Ldate|log.Ltime|log.Lshortfile),
		PanicLogger: log.New(os.Stderr, "PANIC: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *defaultLogger) Print(args ...interface{}) {
	if l.enabled {
		l.InfoLogger.Print(args)
	}
}

func (l *defaultLogger) Printf(arg0 string, args ...interface{}) {
	if l.enabled {
		l.InfoLogger.Printf(arg0, args)
	}
}

func (l *defaultLogger) Println(args ...interface{}) {
	if l.enabled {
		l.InfoLogger.Println(args)
	}
}

func (l *defaultLogger) Fatal(args ...interface{}) {
	if l.enabled {
		l.FatalLogger.Fatal(args)
	}
}
func (l *defaultLogger) Fatalf(arg0 string, args ...interface{}) {
	if l.enabled {
		l.FatalLogger.Fatalf(arg0, args)
	}
}

func (l *defaultLogger) Fatalln(args ...interface{}) {
	if l.enabled {
		l.FatalLogger.Fatalln(args)
	}
}

func (l *defaultLogger) Panic(args ...interface{}) {
	if l.enabled {
		l.PanicLogger.Panic(args)
	}
}

func (l *defaultLogger) Panicf(args0 string, args ...interface{}) {
	if l.enabled {
		l.PanicLogger.Panicf(args0, args)
	}
}

func (l *defaultLogger) Panicln(args ...interface{}) {
	if l.enabled {
		l.PanicLogger.Panicln(args)
	}
}
