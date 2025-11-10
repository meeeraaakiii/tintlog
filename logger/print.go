// Wrappers arond Log and LogBool
package tl

import (
	"strings"

	"github.com/meeeraaakiii/tintlog/palette"
)

// LogJSON prints a labeled JSON block using the standard format.
// Example:
//   LogJSON(tl.Info, palette.CyanDim, "description", value)
func LogJSON(level LogLevel, colorize palette.Colorizer, title string, value string) {
	Log(level, colorize, "%s (JSON):\n'''\n%s\n'''", title, value)
}

// LogRewrite writes a line prefixed with \r and WITHOUT a trailing newline,
// allowing you to re-draw the same terminal line (e.g., progress updates).
// Note: the carriage return is placed at the start of the message body;
// ts/prefix still print normally, so the body portion is what's “rewritten”.
func LogRewrite(level LogLevel, colorize palette.Colorizer, format string, args ...any) {
	LogBool(level, colorize, false, "\r"+format+strings.Repeat(" ", 20), args...)
}
