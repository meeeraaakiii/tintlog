// logger/logger.go
package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

var (
	LoggerOutput      io.Writer = os.Stderr
	LoggerOutputMutex sync.Mutex
)

// Colorizer post-processes a fully formatted message (handles multiline safely).
type Colorizer func(string) string

// Log formats and writes the message, coloring only the args via colorize.
func Log(colorize Colorizer, format string, args ...any) {
	// Colorize arguments only (strings, errors, and fmt.Stringer)
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
				// leave non-strings (numbers, structs, etc.) untouched
			}
		}
	}

	msg := fmt.Sprintf(format, args...)
	// optional: ensure each log call ends with a newline
	if !strings.HasSuffix(msg, "\n") {
		msg += "\n"
	}

	LoggerOutputMutex.Lock()
	_, _ = io.WriteString(LoggerOutput, msg)
	LoggerOutputMutex.Unlock()
}
