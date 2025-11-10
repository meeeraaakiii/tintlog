package tl

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/meeeraaakiii/go-tintlog/color"
)

var (
	LoggerFile      *os.File
	LoggerFilePath  string
	LoggerFileMutex sync.Mutex
)

type LogLine struct {
	Time   time.Time `json:"time"`
	TID    int       `json:"tid,omitempty"`
	Level  LogLevel  `json:"level"`
	Color  string    `json:"color,omitempty"` // e.g., "Green"
	Format string    `json:"format"`          // original format string
	Args   []any     `json:"args"`            // sanitized args (no ANSI)
}

/*
This will initiate json logging file with name like Cfg.LogFileFormat and contents like those:
{"t": timestamp, "tid": yourthreadid, "l": loglevel, "msg": "log message with \033[0;32mcolor\033[0m"}
{"t": timestamp, "tid": yourthreadid, "l": loglevel, "msg": "another log message with \033[0;32mcolor\033[0m"}

This function is only called if an option to save to logger file is specified when initializing logr.

This function will change Cfg.LoggerFilePath and Cfg.LoggerFile
*/
func OpenLoggerFile(logDir string) (err error, errMsg string) {
	err, errMsg = CreateDirIfDoesntExist(logDir)
	if err != nil {
		return err, errMsg
	}

	LoggerFilePath = filepath.Join(logDir, time.Now().Format(Cfg.LogFileFormat))

	Log(Notice, color.Blue, "%s log file '%s'", "Creating", LoggerFilePath)
	LoggerFile, err = os.OpenFile(LoggerFilePath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err, fmt.Sprintf("Unable to open file: '%s'", LoggerFilePath)
	}

	Log(Notice1, color.Green, "%s log file '%v'", "Created", LoggerFilePath)
	return nil, ""
}

// Cfg.LoggerFilePath == "" at the point of logger initialization, so
// it will just print without saving to log file
func CreateDirIfDoesntExist(path string) (err error, errMsg string) {
	Log(Info, color.Blue, "%s dir: '%s'", "Creating", path)
	if path == "" {
		Log(Verbose2, color.DimPurple, "%s", "Dir is an empty string, not creating")
		return nil, ""
	}
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		Log(Verbose, color.DimPurple, "Dir '%s' doesn't exist", path)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err, fmt.Sprintf("Unable to create dir: %s", path)
		}
	} else {
		Log(Info1, color.DimPurple, "Dir '%s' already exists, not creating", path)
		return nil, ""
	}

	Log(Info1, color.Green, "%s dir: '%s'", "Creating", path)

	return nil, ""
}

func writeLogJSONL(line LogLine) {
	if LoggerFilePath == "" {
		return
	}
	b, err := json.Marshal(line)
	if err != nil {
		return
	}
	LoggerFileMutex.Lock()
	_, _ = LoggerFile.Write(append(b, '\n'))
	LoggerFileMutex.Unlock()
}
