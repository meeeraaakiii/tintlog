package palette

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
	"RedBright":    FgColorizer("RedBright", BrightRedColor, false),
	"OrangeBright": FgColorizer("OrangeBright", BrightOrangeColor, false),
	"YellowBright": FgColorizer("YellowBright", BrightYellowColor, false),
	"GreenBright":  FgColorizer("GreenBright", BrightGreenColor, false),
	"CyanBright":   FgColorizer("CyanBright", BrightCyanColor, false),
	"BlueBright":   FgColorizer("BlueBright", BrightBlueColor, false),
	"PurpleBright": FgColorizer("PurpleBright", BrightPurpleColor, false),
	"GrayBright":   FgColorizer("GrayBright", BrightGrayColor, false),

	// Dim shades
	"RedDim":    FgColorizer("RedDim", DimRedColor, false),
	"OrangeDim": FgColorizer("OrangeDim", DimOrangeColor, false),
	"YellowDim": FgColorizer("YellowDim", DimYellowColor, false),
	"GreenDim":  FgColorizer("GreenDim", DimGreenColor, false),
	"CyanDim":   FgColorizer("CyanDim", DimCyanColor, false),
	"BlueDim":   FgColorizer("BlueDim", DimBlueColor, false),
	"PurpleDim": FgColorizer("PurpleDim", DimPurpleColor, false),
	"GrayDim":   FgColorizer("GrayDim", DimGrayColor, false),

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
	"RedBrightBold":    FgColorizer("RedBrightBold", BrightRedColor, true),
	"OrangeBrightBold": FgColorizer("OrangeBrightBold", BrightOrangeColor, true),
	"YellowBrightBold": FgColorizer("YellowBrightBold", BrightYellowColor, true),
	"GreenBrightBold":  FgColorizer("GreenBrightBold", BrightGreenColor, true),
	"CyanBrightBold":   FgColorizer("CyanBrightBold", BrightCyanColor, true),
	"BlueBrightBold":   FgColorizer("BlueBrightBold", BrightBlueColor, true),
	"PurpleBrightBold": FgColorizer("PurpleBrightBold", BrightPurpleColor, true),
	"GrayBrightBold":   FgColorizer("GrayBrightBold", BrightGrayColor, true),

	// --- Bold counterparts (Dim shades) ---
	"RedDimBold":    FgColorizer("RedDimBold", DimRedColor, true),
	"OrangeDimBold": FgColorizer("OrangeDimBold", DimOrangeColor, true),
	"YellowDimBold": FgColorizer("YellowDimBold", DimYellowColor, true),
	"GreenDimBold":  FgColorizer("GreenDimBold", DimGreenColor, true),
	"CyanDimBold":   FgColorizer("CyanDimBold", DimCyanColor, true),
	"BlueDimBold":   FgColorizer("BlueDimBold", DimBlueColor, true),
	"PurpleDimBold": FgColorizer("PurpleDimBold", DimPurpleColor, true),
	"GrayDimBold":   FgColorizer("GrayDimBold", DimGrayColor, true),

	// --- Background variants (Base hues; black fg on color bg) ---
	"RedBackground":    FgBgColorizer("RedBackground", BlackColor, RedColor, false),
	"OrangeBackground": FgBgColorizer("OrangeBackground", BlackColor, OrangeColor, false),
	"YellowBackground": FgBgColorizer("YellowBackground", BlackColor, YellowColor, false),
	"GreenBackground":  FgBgColorizer("GreenBackground", BlackColor, GreenColor, false),
	"CyanBackground":   FgBgColorizer("CyanBackground", BlackColor, CyanColor, false),
	"BlueBackground":   FgBgColorizer("BlueBackground", BlackColor, BlueColor, false),
	"PurpleBackground": FgBgColorizer("PurpleBackground", BlackColor, PurpleColor, false),
	"GrayBackground":   FgBgColorizer("GrayBackground", BlackColor, GrayColor, false),

	// --- Bold background variants (Base hues; black fg on color bg) ---
	"RedBoldBackground":    FgBgColorizer("RedBoldBackground", BlackColor, RedColor, true),
	"OrangeBoldBackground": FgBgColorizer("OrangeBoldBackground", BlackColor, OrangeColor, true),
	"YellowBoldBackground": FgBgColorizer("YellowBoldBackground", BlackColor, YellowColor, true),
	"GreenBoldBackground":  FgBgColorizer("GreenBoldBackground", BlackColor, GreenColor, true),
	"CyanBoldBackground":   FgBgColorizer("CyanBoldBackground", BlackColor, CyanColor, true),
	"BlueBoldBackground":   FgBgColorizer("BlueBoldBackground", BlackColor, BlueColor, true),
	"PurpleBoldBackground": FgBgColorizer("PurpleBoldBackground", BlackColor, PurpleColor, true),
	"GrayBoldBackground":   FgBgColorizer("GrayBoldBackground", BlackColor, GrayColor, true),
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
	RedBright    = Colorizers["RedBright"]
	OrangeBright = Colorizers["OrangeBright"]
	YellowBright = Colorizers["YellowBright"]
	GreenBright  = Colorizers["GreenBright"]
	CyanBright   = Colorizers["CyanBright"]
	BlueBright   = Colorizers["BlueBright"]
	PurpleBright = Colorizers["PurpleBright"]
	GrayBright   = Colorizers["GrayBright"]

	// Dim shades
	RedDim    = Colorizers["RedDim"]
	OrangeDim = Colorizers["OrangeDim"]
	YellowDim = Colorizers["YellowDim"]
	GreenDim  = Colorizers["GreenDim"]
	CyanDim   = Colorizers["CyanDim"]
	BlueDim   = Colorizers["BlueDim"]
	PurpleDim = Colorizers["PurpleDim"]
	GrayDim   = Colorizers["GrayDim"]

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
	RedBrightBold    = Colorizers["RedBrightBold"]
	OrangeBrightBold = Colorizers["OrangeBrightBold"]
	YellowBrightBold = Colorizers["YellowBrightBold"]
	GreenBrightBold  = Colorizers["GreenBrightBold"]
	CyanBrightBold   = Colorizers["CyanBrightBold"]
	BlueBrightBold   = Colorizers["BlueBrightBold"]
	PurpleBrightBold = Colorizers["PurpleBrightBold"]
	GrayBrightBold   = Colorizers["GrayBrightBold"]

	// Bold counterparts (Dim)
	RedDimBold    = Colorizers["RedDimBold"]
	OrangeDimBold = Colorizers["OrangeDimBold"]
	YellowDimBold = Colorizers["YellowDimBold"]
	GreenDimBold  = Colorizers["GreenDimBold"]
	CyanDimBold   = Colorizers["CyanDimBold"]
	BlueDimBold   = Colorizers["BlueDimBold"]
	PurpleDimBold = Colorizers["PurpleDimBold"]
	GrayDimBold   = Colorizers["GrayDimBold"]

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
