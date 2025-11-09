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
	"Red":    FgColorizer("Red", Red, false),
	"Orange": FgColorizer("Orange", Orange, false),
	"Yellow": FgColorizer("Yellow", Yellow, false),
	"Green":  FgColorizer("Green", Green, false),
	"Cyan":   FgColorizer("Cyan", Cyan, false),
	"Blue":   FgColorizer("Blue", Blue, false),
	"Purple": FgColorizer("Purple", Purple, false),
	"Gray":   FgColorizer("Gray", Gray, false),

	// Bright tints
	"BrightRed":    FgColorizer("BrightRed", BrightRed, false),
	"BrightOrange": FgColorizer("BrightOrange", BrightOrange, false),
	"BrightYellow": FgColorizer("BrightYellow", BrightYellow, false),
	"BrightGreen":  FgColorizer("BrightGreen", BrightGreen, false),
	"BrightCyan":   FgColorizer("BrightCyan", BrightCyan, false),
	"BrightBlue":   FgColorizer("BrightBlue", BrightBlue, false),
	"BrightPurple": FgColorizer("BrightPurple", BrightPurple, false),
	"BrightGray":   FgColorizer("BrightGray", BrightGray, false),

	// Dim shades
	"DimRed":    FgColorizer("DimRed", DimRed, false),
	"DimOrange": FgColorizer("DimOrange", DimOrange, false),
	"DimYellow": FgColorizer("DimYellow", DimYellow, false),
	"DimGreen":  FgColorizer("DimGreen", DimGreen, false),
	"DimCyan":   FgColorizer("DimCyan", DimCyan, false),
	"DimBlue":   FgColorizer("DimBlue", DimBlue, false),
	"DimPurple": FgColorizer("DimPurple", DimPurple, false),
	"DimGray":   FgColorizer("DimGray", DimGray, false),

	// No color
	"NoColor": {Name: "NoColor", Fn: nil},

	// Example bold presets (add more if you like)
	"RedBold":           FgColorizer("RedBold", Red, true),
	"GreenBold":         FgColorizer("GreenBold", Green, true),
	"BlueBold":          FgColorizer("BlueBold", Blue, true),
	"RedBackground":     FgBgColorizer("RedBackground", Black, Red, false),
	"RedBoldBackground": FgBgColorizer("RedBoldBackground", Black, Red, true),
}

/* --------------- convenience aliases (import-friendly) ----------------- */

var (
	RedText    = Colorizers["Red"]
	OrangeText = Colorizers["Orange"]
	YellowText = Colorizers["Yellow"]
	GreenText  = Colorizers["Green"]
	CyanText   = Colorizers["Cyan"]
	BlueText   = Colorizers["Blue"]
	PurpleText = Colorizers["Purple"]
	GrayText   = Colorizers["Gray"]

	BrightRedText    = Colorizers["BrightRed"]
	BrightOrangeText = Colorizers["BrightOrange"]
	BrightYellowText = Colorizers["BrightYellow"]
	BrightGreenText  = Colorizers["BrightGreen"]
	BrightCyanText   = Colorizers["BrightCyan"]
	BrightBlueText   = Colorizers["BrightBlue"]
	BrightPurpleText = Colorizers["BrightPurple"]
	BrightGrayText   = Colorizers["BrightGray"]

	DimRedText    = Colorizers["DimRed"]
	DimOrangeText = Colorizers["DimOrange"]
	DimYellowText = Colorizers["DimYellow"]
	DimGreenText  = Colorizers["DimGreen"]
	DimCyanText   = Colorizers["DimCyan"]
	DimBlueText   = Colorizers["DimBlue"]
	DimPurpleText = Colorizers["DimPurple"]
	DimGrayText   = Colorizers["DimGray"]

	NoColor = Colorizers["NoColor"]

	// Bold examples
	RedBoldText   = Colorizers["RedBold"]
	GreenBoldText = Colorizers["GreenBold"]
	BlueBoldText  = Colorizers["BlueBold"]
	RedBackgroundText = Colorizers["RedBackground"]
	RedBoldBackgroundText = Colorizers["RedBoldBackground"]
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
