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
	out io.Writer = os.Stdout
	mu  sync.Mutex
)

// Colorizer post-processes a fully formatted message (handles multiline safely).
type Colorizer func(string) string

// SetOutput lets you redirect logs (files, buffers, etc.).
func SetOutput(w io.Writer) { if w != nil { out = w } }

// Log formats and writes the message, coloring only the args via colorize.
// Example:
//   logger.Log(logger.Color, "Regular color text %s", "text printed with logger.Color\nstillsame logger.Color")
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

	mu.Lock()
	_, _ = io.WriteString(out, msg)
	mu.Unlock()
}
