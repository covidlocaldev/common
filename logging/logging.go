package logging

import (
	"fmt"
	"log"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

func (a LogLevel) String() string {
	return [...]string{"DEBUG", "INFO", "WARN", "ERROR"}[a]
}

func Debug(format string, v... interface{}) {
	printLog(DEBUG, format, v)
}

func Info(format string, v... interface{}) {
	printLog(INFO, format, v)
}

func Warn(format string, v... interface{}) {
	printLog(WARN, format, v)
}

func Error(format string, v... interface{}) {
	printLog(ERROR, format, v)
}

func printLog(level LogLevel, format string, v... interface{}) {
	message := fmt.Sprintf(format, v...)

	log.Printf("[%s] %s", level.String(), message)
}
