package file

import (
	"log"
	"os"
)

type (
	fileLogger struct {
		Log *log.Logger
	}
)

// FileLogger is a logger that satisfies Logger interface and depending on environment may put logs into a file
var FileLogger fileLogger

func init() {
	f, err := os.OpenFile("testlogfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	//defer f.Close()

	FileLogger = fileLogger{Log: log.New(f, "", log.LstdFlags|log.Lshortfile)}
}

func (f fileLogger) Errorf(format string, args ...interface{}) {
	f.Log.Printf(format, args...)
}
