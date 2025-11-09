// logger/color.go
package logger

import (
	"fmt"
	"strings"
)

const reset = "\x1b[0m"

type RGB struct{ R, G, B uint8 }

// ----- low-level (single line) -----

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

// ----- multiline-safe (per-line wrap) -----

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

// ----- presets & ready-to-use Colorizers -----

var (
	Red   = RGB{255, 0, 0}
	Green = RGB{0, 255, 0}
	Blue  = RGB{0, 0, 255}

	SoftYellowBG = RGB{0xFF, 0xF8, 0xE1}
	SoftGreenBG  = RGB{0xEC, 0xFD, 0xF5}
)

// Public sentinel for clarity.

// Colorizers you can pass to Log
var (
	Color        Colorizer = func(s string) string { return FgLines(s, Green) } // default demo color
	RedText      Colorizer = func(s string) string { return FgLines(s, Red) }
	GreenText    Colorizer = func(s string) string { return FgLines(s, Green) }
	BlueText     Colorizer = func(s string) string { return FgLines(s, Blue) }
	OnSoftYellow Colorizer = func(s string) string { return FgBgLines(s, RGB{0x43, 0x62, 0x12}, SoftYellowBG) }
	OnSoftGreen  Colorizer = func(s string) string { return FgBgLines(s, RGB{0x16, 0x65, 0x34}, SoftGreenBG) }
	NoColor      Colorizer = nil
)
