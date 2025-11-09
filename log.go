// logger/logger.go
package logger

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	LoggerOutput      io.Writer = os.Stderr
	LoggerOutputMutex sync.Mutex
)

// Colorizer is assumed to be your existing type, e.g.:
// type Colorizer func(s string) string

// LogAt prints with an explicit log level and optional goroutine TID (if enabled in Cfg.UseTid).
// It respects Cfg.LogLevel (only prints when Cfg.LogLevel >= level).
// It colorizes:
//   • the [Level] (and [tid] if present) using the provided colorize
//   • string-ish args (string, error, fmt.Stringer) inside the message body
func Log(level LogLevel, colorize Colorizer, format string, args ...any) {
	// respect log level threshold; file duplication stays out of scope for now
	if Cfg.LogLevel < level {
		return
	}

	// colorize arguments only (strings, errors, fmt.Stringer)
	if colorize != nil {
		for i, a := range args {
			switch v := a.(type) {
			case string:
				args[i] = colorize(v)
			case error:
				args[i] = colorize(v.Error())
			case fmt.Stringer:
				args[i] = colorize(v.String())
			default:
				// leave non-strings untouched
			}
		}
	}

	body := fmt.Sprintf(format, args...)
	if !strings.HasSuffix(body, "\n") {
		body += "\n"
	}

	// build prefix: [Level] or [Level][tid]
	levelStr := level.String()
	if colorize != nil {
		levelStr = colorize(levelStr)
	}
	prefix := "[" + levelStr + "] "

	if Cfg.UseTid != nil && *Cfg.UseTid {
		tid := getTid() // you already have this in your old codebase
		tidStr := strconv.Itoa(tid)
		if colorize != nil {
			tidStr = colorize(tidStr)
		}
		prefix = "[" + levelStr + "][" + tidStr + "] "
	}

	LoggerOutputMutex.Lock()
	_, _ = io.WriteString(LoggerOutput, prefix+body)
	LoggerOutputMutex.Unlock()
}

