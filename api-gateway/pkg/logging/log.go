package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setprefix(DEBUG)
	logger.Println(v...)
}

func Info(v ...interface{}) {
	setprefix(INFO)
	logger.Println(v...)
}

func Warn(v ...interface{}) {
	setprefix(WARNING)
	logger.Println(v...)
}

func Error(v ...interface{}) {
	setprefix(ERROR)
	logger.Println(v...)
}

func Fatal(v ...interface{}) {
	setprefix(FATAL)
	logger.Fatalln(v...)
}

func setprefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
