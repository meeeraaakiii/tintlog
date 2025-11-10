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

	// --- Bold counterparts (Base hues) ---
	"RedBold":    FgColorizer("RedBold", RedColor, true),
	"OrangeBold": FgColorizer("OrangeBold", OrangeColor, true),
	"YellowBold": FgColorizer("YellowBold", YellowColor, true),
	"GreenBold":  FgColorizer("GreenBold", GreenColor, true),
	"CyanBold":   FgColorizer("CyanBold", CyanColor, true),
	"BlueBold":   FgColorizer("BlueBold", BlueColor, true),
	"PurpleBold": FgColorizer("PurpleBold", PurpleColor, true),
	"GrayBold":   FgColorizer("GrayBold", GrayColor, true),

	// --- Bold counterparts (Bright tints) ---
	"BrightRedBold":    FgColorizer("BrightRedBold", BrightRedColor, true),
	"BrightOrangeBold": FgColorizer("BrightOrangeBold", BrightOrangeColor, true),
	"BrightYellowBold": FgColorizer("BrightYellowBold", BrightYellowColor, true),
	"BrightGreenBold":  FgColorizer("BrightGreenBold", BrightGreenColor, true),
	"BrightCyanBold":   FgColorizer("BrightCyanBold", BrightCyanColor, true),
	"BrightBlueBold":   FgColorizer("BrightBlueBold", BrightBlueColor, true),
	"BrightPurpleBold": FgColorizer("BrightPurpleBold", BrightPurpleColor, true),
	"BrightGrayBold":   FgColorizer("BrightGrayBold", BrightGrayColor, true),

	// --- Bold counterparts (Dim shades) ---
	"DimRedBold":    FgColorizer("DimRedBold", DimRedColor, true),
	"DimOrangeBold": FgColorizer("DimOrangeBold", DimOrangeColor, true),
	"DimYellowBold": FgColorizer("DimYellowBold", DimYellowColor, true),
	"DimGreenBold":  FgColorizer("DimGreenBold", DimGreenColor, true),
	"DimCyanBold":   FgColorizer("DimCyanBold", DimCyanColor, true),
	"DimBlueBold":   FgColorizer("DimBlueBold", DimBlueColor, true),
	"DimPurpleBold": FgColorizer("DimPurpleBold", DimPurpleColor, true),
	"DimGrayBold":   FgColorizer("DimGrayBold", DimGrayColor, true),

	// --- Background variants (Base hues; black fg on color bg) ---
	"RedBackground":     FgBgColorizer("RedBackground", BlackColor, RedColor, false),
	"OrangeBackground":  FgBgColorizer("OrangeBackground", BlackColor, OrangeColor, false),
	"YellowBackground":  FgBgColorizer("YellowBackground", BlackColor, YellowColor, false),
	"GreenBackground":   FgBgColorizer("GreenBackground", BlackColor, GreenColor, false),
	"CyanBackground":    FgBgColorizer("CyanBackground", BlackColor, CyanColor, false),
	"BlueBackground":    FgBgColorizer("BlueBackground", BlackColor, BlueColor, false),
	"PurpleBackground":  FgBgColorizer("PurpleBackground", BlackColor, PurpleColor, false),
	"GrayBackground":    FgBgColorizer("GrayBackground", BlackColor, GrayColor, false),

	// --- Bold background variants (Base hues; black fg on color bg) ---
	"RedBoldBackground":     FgBgColorizer("RedBoldBackground", BlackColor, RedColor, true),
	"OrangeBoldBackground":  FgBgColorizer("OrangeBoldBackground", BlackColor, OrangeColor, true),
	"YellowBoldBackground":  FgBgColorizer("YellowBoldBackground", BlackColor, YellowColor, true),
	"GreenBoldBackground":   FgBgColorizer("GreenBoldBackground", BlackColor, GreenColor, true),
	"CyanBoldBackground":    FgBgColorizer("CyanBoldBackground", BlackColor, CyanColor, true),
	"BlueBoldBackground":    FgBgColorizer("BlueBoldBackground", BlackColor, BlueColor, true),
	"PurpleBoldBackground":  FgBgColorizer("PurpleBoldBackground", BlackColor, PurpleColor, true),
	"GrayBoldBackground":    FgBgColorizer("GrayBoldBackground", BlackColor, GrayColor, true),
}

/* --------------- convenience aliases (import-friendly) ----------------- */
var (
	// Base hues
	Red    = Colorizers["Red"]
	Orange = Colorizers["Orange"]
	Yellow = Colorizers["Yellow"]
	Green  = Colorizers["Green"]
	Cyan   = Colorizers["Cyan"]
	Blue   = Colorizers["Blue"]
	Purple = Colorizers["Purple"]
	Gray   = Colorizers["Gray"]

	// Bright tints
	BrightRed    = Colorizers["BrightRed"]
	BrightOrange = Colorizers["BrightOrange"]
	BrightYellow = Colorizers["BrightYellow"]
	BrightGreen  = Colorizers["BrightGreen"]
	BrightCyan   = Colorizers["BrightCyan"]
	BrightBlue   = Colorizers["BrightBlue"]
	BrightPurple = Colorizers["BrightPurple"]
	BrightGray   = Colorizers["BrightGray"]

	// Dim shades
	DimRed    = Colorizers["DimRed"]
	DimOrange = Colorizers["DimOrange"]
	DimYellow = Colorizers["DimYellow"]
	DimGreen  = Colorizers["DimGreen"]
	DimCyan   = Colorizers["DimCyan"]
	DimBlue   = Colorizers["DimBlue"]
	DimPurple = Colorizers["DimPurple"]
	DimGray   = Colorizers["DimGray"]

	// No color
	NoColor = Colorizers["NoColor"]

	// Bold counterparts (Base)
	RedBold    = Colorizers["RedBold"]
	OrangeBold = Colorizers["OrangeBold"]
	YellowBold = Colorizers["YellowBold"]
	GreenBold  = Colorizers["GreenBold"]
	CyanBold   = Colorizers["CyanBold"]
	BlueBold   = Colorizers["BlueBold"]
	PurpleBold = Colorizers["PurpleBold"]
	GrayBold   = Colorizers["GrayBold"]

	// Bold counterparts (Bright)
	BrightRedBold    = Colorizers["BrightRedBold"]
	BrightOrangeBold = Colorizers["BrightOrangeBold"]
	BrightYellowBold = Colorizers["BrightYellowBold"]
	BrightGreenBold  = Colorizers["BrightGreenBold"]
	BrightCyanBold   = Colorizers["BrightCyanBold"]
	BrightBlueBold   = Colorizers["BrightBlueBold"]
	BrightPurpleBold = Colorizers["BrightPurpleBold"]
	BrightGrayBold   = Colorizers["BrightGrayBold"]

	// Bold counterparts (Dim)
	DimRedBold    = Colorizers["DimRedBold"]
	DimOrangeBold = Colorizers["DimOrangeBold"]
	DimYellowBold = Colorizers["DimYellowBold"]
	DimGreenBold  = Colorizers["DimGreenBold"]
	DimCyanBold   = Colorizers["DimCyanBold"]
	DimBlueBold   = Colorizers["DimBlueBold"]
	DimPurpleBold = Colorizers["DimPurpleBold"]
	DimGrayBold   = Colorizers["DimGrayBold"]

	// Background variants (Base; black fg on color bg)
	RedBackground    = Colorizers["RedBackground"]
	OrangeBackground = Colorizers["OrangeBackground"]
	YellowBackground = Colorizers["YellowBackground"]
	GreenBackground  = Colorizers["GreenBackground"]
	CyanBackground   = Colorizers["CyanBackground"]
	BlueBackground   = Colorizers["BlueBackground"]
	PurpleBackground = Colorizers["PurpleBackground"]
	GrayBackground   = Colorizers["GrayBackground"]

	// Bold background variants (Base; black fg on color bg)
	RedBoldBackground    = Colorizers["RedBoldBackground"]
	OrangeBoldBackground = Colorizers["OrangeBoldBackground"]
	YellowBoldBackground = Colorizers["YellowBoldBackground"]
	GreenBoldBackground  = Colorizers["GreenBoldBackground"]
	CyanBoldBackground   = Colorizers["CyanBoldBackground"]
	BlueBoldBackground   = Colorizers["BlueBoldBackground"]
	PurpleBoldBackground = Colorizers["PurpleBoldBackground"]
	GrayBoldBackground   = Colorizers["GrayBoldBackground"]
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
