package color

import (
	"fmt"
	"strings"
)

// Colorizer applies ANSI color to a string (per line). Nil Fn = no color.
type Colorizer struct {
	Name string
	Fn   func(string) string
}

// Apply safely applies the colorizer if present.
func (c Colorizer) Apply(s string) string {
	if c.Fn == nil {
		return s
	}
	return c.Fn(s)
}

/* ---------- internal helpers for bold/non-bold line coloring ---------- */

func fgLines(s string, fg Color, bold bool) string {
	lines, trail := splitKeepTrail(s)
	r := fg.MustRGB()
	var prefix string
	if bold {
		prefix = fmt.Sprintf("\x1b[1;38;2;%d;%d;%dm", r.R, r.G, r.B)
	} else {
		prefix = fmt.Sprintf("\x1b[38;2;%d;%d;%dm", r.R, r.G, r.B)
	}
	for i, ln := range lines {
		lines[i] = prefix + ln + reset
	}
	return strings.Join(lines, "\n") + trail
}

func fgBgLines(s string, fg, bg Color, bold bool) string {
	lines, trail := splitKeepTrail(s)
	fr := fg.MustRGB()
	br := bg.MustRGB()
	var prefix string
	if bold {
		prefix = fmt.Sprintf("\x1b[1;38;2;%d;%d;%d;48;2;%d;%d;%dm", fr.R, fr.G, fr.B, br.R, br.G, br.B)
	} else {
		prefix = fmt.Sprintf("\x1b[38;2;%d;%d;%d;48;2;%d;%d;%dm", fr.R, fr.G, fr.B, br.R, br.G, br.B)
	}
	for i, ln := range lines {
		lines[i] = prefix + ln + reset
	}
	return strings.Join(lines, "\n") + trail
}

/* -------------------- builders to create colorizers -------------------- */

// FgColorizer returns a per-line foreground colorizer; set bold=true for bold text.
func FgColorizer(name string, fg Color, bold bool) Colorizer {
	return Colorizer{
		Name: name,
		Fn:   func(s string) string { return fgLines(s, fg, bold) },
	}
}

// FgBgColorizer returns a per-line foreground+background colorizer; set bold=true for bold text.
func FgBgColorizer(name string, fg, bg Color, bold bool) Colorizer {
	return Colorizer{
		Name: name,
		Fn:   func(s string) string { return fgBgLines(s, fg, bg, bold) },
	}
}

/* ---------------- registry of reusable colorizers (non-bold) ----------- */

var Colorizers = map[string]Colorizer{
	// Base hues
	"Red":    FgColorizer("Red", RedColor, false),
	"Orange": FgColorizer("Orange", OrangeColor, false),
	"Yellow": FgColorizer("Yellow", YellowColor, false),
	"Green":  FgColorizer("Green", GreenColor, false),
	"Cyan":   FgColorizer("Cyan", CyanColor, false),
	"Blue":   FgColorizer("Blue", BlueColor, false),
	"Purple": FgColorizer("Purple", PurpleColor, false),
	"Gray":   FgColorizer("Gray", GrayColor, false),

	// Bright tints
	"BrightRed":    FgColorizer("BrightRed", BrightRedColor, false),
	"BrightOrange": FgColorizer("BrightOrange", BrightOrangeColor, false),
	"BrightYellow": FgColorizer("BrightYellow", BrightYellowColor, false),
	"BrightGreen":  FgColorizer("BrightGreen", BrightGreenColor, false),
	"BrightCyan":   FgColorizer("BrightCyan", BrightCyanColor, false),
	"BrightBlue":   FgColorizer("BrightBlue", BrightBlueColor, false),
	"BrightPurple": FgColorizer("BrightPurple", BrightPurpleColor, false),
	"BrightGray":   FgColorizer("BrightGray", BrightGrayColor, false),

	// Dim shades
	"DimRed":    FgColorizer("DimRed", DimRedColor, false),
	"DimOrange": FgColorizer("DimOrange", DimOrangeColor, false),
	"DimYellow": FgColorizer("DimYellow", DimYellowColor, false),
	"DimGreen":  FgColorizer("DimGreen", DimGreenColor, false),
	"DimCyan":   FgColorizer("DimCyan", DimCyanColor, false),
	"DimBlue":   FgColorizer("DimBlue", DimBlueColor, false),
	"DimPurple": FgColorizer("DimPurple", DimPurpleColor, false),
	"DimGray":   FgColorizer("DimGray", DimGrayColor, false),

	// No color
	"NoColor": {Name: "NoColor", Fn: nil},

	// Example bold presets (add more if you like)
	"RedBold":           FgColorizer("RedBold", RedColor, true),
	"GreenBold":         FgColorizer("GreenBold", GreenColor, true),
	"BlueBold":          FgColorizer("BlueBold", BlueColor, true),
	"RedBackground":     FgBgColorizer("RedBackground", BlackColor, RedColor, false),
	"RedBoldBackground": FgBgColorizer("RedBoldBackground", BlackColor, RedColor, true),
}

/* --------------- convenience aliases (import-friendly) ----------------- */

var (
	Red    = Colorizers["Red"]
	Orange = Colorizers["Orange"]
	Yellow = Colorizers["Yellow"]
	Green  = Colorizers["Green"]
	Cyan   = Colorizers["Cyan"]
	Blue   = Colorizers["Blue"]
	Purple = Colorizers["Purple"]
	Gray   = Colorizers["Gray"]

	BrightRed    = Colorizers["BrightRed"]
	BrightOrange = Colorizers["BrightOrange"]
	BrightYellow = Colorizers["BrightYellow"]
	BrightGreen  = Colorizers["BrightGreen"]
	BrightCyan   = Colorizers["BrightCyan"]
	BrightBlue   = Colorizers["BrightBlue"]
	BrightPurple = Colorizers["BrightPurple"]
	BrightGray   = Colorizers["BrightGray"]

	DimRed    = Colorizers["DimRed"]
	DimOrange = Colorizers["DimOrange"]
	DimYellow = Colorizers["DimYellow"]
	DimGreen  = Colorizers["DimGreen"]
	DimCyan   = Colorizers["DimCyan"]
	DimBlue   = Colorizers["DimBlue"]
	DimPurple = Colorizers["DimPurple"]
	DimGray   = Colorizers["DimGray"]

	NoColor = Colorizers["NoColor"]

	// Bold examples
	RedBold   = Colorizers["RedBold"]
	GreenBold = Colorizers["GreenBold"]
	BlueBold  = Colorizers["BlueBold"]
	RedBackground = Colorizers["RedBackground"]
	RedBoldBackground = Colorizers["RedBoldBackground"]
)

/* ------------------- registration convenience funcs ------------------- */

// RegisterColorizer adds/overwrites a colorizer in the registry.
func RegisterColorizer(name string, fn func(string) string) Colorizer {
	c := Colorizer{Name: name, Fn: fn}
	Colorizers[name] = c
	return c
}

// RegisterFg registers a foreground-only colorizer; set bold=true for bold text.
func RegisterFg(name string, fg Color, bold bool) Colorizer {
	return RegisterColorizer(name, func(s string) string { return fgLines(s, fg, bold) })
}

// RegisterFgBg registers a fg+bg colorizer; set bold=true for bold text.
func RegisterFgBg(name string, fg, bg Color, bold bool) Colorizer {
	return RegisterColorizer(name, func(s string) string { return fgBgLines(s, fg, bg, bold) })
}

// splitKeepTrail splits s by '\n' and preserves a single trailing newline (LF or CRLF) if present.
// Returns the lines (without the trailing newline) and the exact trailing newline sequence.
func splitKeepTrail(s string) (lines []string, trailing string) {
	// Preserve exactly one trailing newline (LF or CRLF)
	if strings.HasSuffix(s, "\r\n") {
		trailing = "\r\n"
		s = strings.TrimSuffix(s, "\r\n")
	} else if strings.HasSuffix(s, "\n") {
		trailing = "\n"
		s = strings.TrimSuffix(s, "\n")
	}
	// Split remaining content on '\n' (handles CRLF already stripped above)
	return strings.Split(s, "\n"), trailing
}
