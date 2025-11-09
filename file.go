package logger

import (
	"os"
	"sync"
)

var (
	LoggerFile *os.File
	LoggerFilePath string
	LoggerFileMutex sync.Mutex
)