package utils

import (
	"os"
	"fmt"
	"time"
)

const (
	LOG_LEVEL_ERROR string = "ERROR"
	LOG_LEVEL_INFO  string = "INFO"
	LOG_LEVEL_DEBUG string = "DEBUG"
	LOG_FILE_NAME   string = "summary.log"
)

type Logger interface {
	Error(message string)
	Info(message string)
	Debug(message string)
}

type FileLogger struct {
	message string
	file    *os.File
	logger FileLogger
}

func (logger *FileLogger) Error(message string) {
	logger.write(message, LOG_LEVEL_ERROR)
}

func (logger *FileLogger) Info(message string) {
	logger.write(message, LOG_LEVEL_INFO)
}

func (logger *FileLogger) Debug(message string) {
	logger.write(message, LOG_LEVEL_DEBUG)
}

func NewFileLogger() *FileLogger {
	logger := FileLogger{}
	//first check file is exist,it not then create it
	if _, err := os.Stat(LOG_FILE_NAME); os.IsNotExist(err) {
		logger.file, err = os.Create(LOG_FILE_NAME)
		checkError(err)
		logger.file.Close()
	}

	return &logger
}

func (logger *FileLogger) write(message string, logLevel string) {
	logger.openFile()

	strFormat := "[%v] [%s] %s"
	message = fmt.Sprintf(strFormat, time.Now(), logLevel, message)
	logger.file.WriteString(message)

	logger.closeFile()
}

func (logger *FileLogger) openFile() {
	file, err := os.OpenFile(LOG_FILE_NAME, os.O_APPEND, 0666)
	checkError(err)

	logger.file = file
}

func (logger *FileLogger) closeFile() {
	logger.file.Close()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
