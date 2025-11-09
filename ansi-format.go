package logger

import (
	"fmt"
	"strings"
)

const reset = "\x1b[0m"

type RGB struct{ R, G, B uint8 }

func Fg(s string, c RGB) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s%s", c.R, c.G, c.B, s, reset)
}
func Bg(s string, c RGB) string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm%s%s", c.R, c.G, c.B, s, reset)
}
func FgBg(s string, fg, bg RGB) string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm\x1b[48;2;%d;%d;%dm%s%s",
		fg.R, fg.G, fg.B, bg.R, bg.G, bg.B, s, reset)
}

func FgLines(s string, c RGB) string {
	lines, trail := splitKeepTrail(s)
	for i, ln := range lines {
		lines[i] = Fg(ln, c)
	}
	return strings.Join(lines, "\n") + trail
}
func FgBgLines(s string, fg, bg RGB) string {
	lines, trail := splitKeepTrail(s)
	for i, ln := range lines {
		lines[i] = FgBg(ln, fg, bg)
	}
	return strings.Join(lines, "\n") + trail
}
func splitKeepTrail(s string) (lines []string, trailing string) {
	if strings.HasSuffix(s, "\n") {
		trailing = "\n"
		s = strings.TrimSuffix(s, "\n")
	}
	return strings.Split(s, "\n"), trailing
}
