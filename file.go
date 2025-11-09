package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	LoggerFile *os.File
	LoggerFilePath string
	LoggerFileMutex sync.Mutex
)


/*
This will initiate json logging file with name like Cfg.LogFileFormat and contents like those:
{"t": timestamp, "tid": yourthreadid, "l": loglevel, "msg": "log message with \033[0;32mcolor\033[0m"}
{"t": timestamp, "tid": yourthreadid, "l": loglevel, "msg": "another log message with \033[0;32mcolor\033[0m"}

This function is only called if an option to save to logger file is specified when initializing logger.

This function will change Cfg.LoggerFilePath and Cfg.LoggerFile
*/
func OpenLoggerFile(logDir string) (err error, errMsg string) {
	err, errMsg = CreateDirIfDoesntExist(logDir)
	if err != nil {
		return err, errMsg
	}

	LoggerFilePath = filepath.Join(logDir, time.Now().Format(Cfg.LogFileFormat))

	Log(Notice, Color, "%s log file '%s'", "Creating", LoggerFilePath)
	LoggerFile, err = os.OpenFile(LoggerFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err, fmt.Sprintf("Unable to open file: '%s'", LoggerFilePath)
	}

	Log(Notice1, Color, "%s log file '%v'", "Created", LoggerFilePath)
	return nil, ""
}
