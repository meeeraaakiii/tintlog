package logger

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// get goroutine id
// gotta have this, it will slow things down but we need to track
// in which goroutine thing happens
func getTid() (tid int) {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	tid, err := strconv.Atoi(idField)
	if err != nil {
		log.Printf("Cannot get goroutine id, Err: %s", err.Error())
		return -1
	}
	return tid
}

type LogLine struct {
	Time  time.Time `json:"time"`
	TID   int       `json:"tid,omitempty"`
	Level LogLevel  `json:"level"`
	Msg   string    `json:"msg"`
}

func ensureLogFileOpen() error {
	LoggerFileMutex.Lock()
	defer LoggerFileMutex.Unlock()

	if LoggerFilePath == "" {
		return nil
	}
	if LoggerFile != nil {
		return nil
	}

	f, err := os.OpenFile(LoggerFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	LoggerFile = f
	return nil
}

func writeLogJSONL(line LogLine) {
	if LoggerFilePath == "" {
		return
	}
	if err := ensureLogFileOpen(); err != nil {
		// don’t panic in the logger; best effort only
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
