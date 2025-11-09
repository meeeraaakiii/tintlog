package color

import (
	"fmt"
	"strings"
)

const reset = "\x1b[0m"

type RGB struct{ R, G, B uint8 }

// Fg wraps s with a 24-bit ANSI foreground color.
func Fg(s string, fg Color) string {
	r := fg.MustRGB()
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s%s", r.R, r.G, r.B, s, reset)
}

// Bg wraps s with a 24-bit ANSI background color.
func Bg(s string, bg Color) string {
	r := bg.MustRGB()
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm%s%s", r.R, r.G, r.B, s, reset)
}

// FgBg wraps s with both foreground and background colors.
func FgBg(s string, fg, bg Color) string {
	fr := fg.MustRGB()
	br := bg.MustRGB()
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm\x1b[48;2;%d;%d;%dm%s%s",
		fr.R, fr.G, fr.B, br.R, br.G, br.B, s, reset)
}

// FgLines applies foreground color to each line independently.
// Preserves a trailing newline if present.
func FgLines(s string, fg Color) string {
	lines, trail := splitLinesKeepNL(s)
	for i, ln := range lines {
		lines[i] = Fg(ln, fg)
	}
	return strings.Join(lines, "\n") + trail
}

// FgBgLines applies fg/bg colors to each line independently.
// Preserves a trailing newline if present.
func FgBgLines(s string, fg, bg Color) string {
	lines, trail := splitLinesKeepNL(s)
	for i, ln := range lines {
		lines[i] = FgBg(ln, fg, bg)
	}
	return strings.Join(lines, "\n") + trail
}

// splitLinesKeepNL splits by '\n' and returns the lines plus a trailing "\n" if it existed.
func splitLinesKeepNL(s string) (lines []string, trailing string) {
	if strings.HasSuffix(s, "\n") {
		trailing = "\n"
		s = strings.TrimSuffix(s, "\n")
	}
	return strings.Split(s, "\n"), trailing
}