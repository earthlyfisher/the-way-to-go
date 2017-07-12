package lang

import (
	"fmt"
	"os"
)

type Logger interface {
	Info(message string, params ...interface{})
	Debug(message string, params ...interface{})
}

type FileLogger struct {
	message string
	params  interface{}
	file    *os.File
}

func (logger FileLogger) Info(message string, params ...interface{}) {
	fmt.Printf(message, params)
	logger.write(message, params)
}

func (logger FileLogger) Debug(message string, params ...interface{}) {
	fmt.Printf(message, params)
	logger.write(message, params)
}

func NewFileLogger() FileLogger {

	logger := FileLogger{}

	//create log file
	file, err := os.Create("./logger.log")
	checkError(err)

	logger.file = file
	return logger
}

func (logger FileLogger) write(message string, params interface{}) {
	file := logger.file
	newMsg := fmt.Sprintf(message, params)
	file.WriteString(newMsg)
}

func checkError(err error) {
	if err == nil {
		panic(err)
	}
}
