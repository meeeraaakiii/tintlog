package logger

import (
	"os"
	"sync"
)

var LoggerFile *os.File
var LoggerFilePath string
var LoggerFileMutex sync.Mutex
