package logger

// Colorizer holds a name and the function that applies color.
// Use *Colorizer so nil means “no color”.
type Colorizer struct {
	Name string
	Fn   func(string) string
}

// Apply safely applies the colorizer if present.
func (c *Colorizer) Apply(s string) string {
	if c == nil || c.Fn == nil {
		return s
	}
	return c.Fn(s)
}

// Registry of reusable colorizers by name.
var Colorizers = map[string]Colorizer{
	"Red":         {Name: "Red", Fn: func(s string) string { return FgLines(s, Red) }},
	"Green":       {Name: "Green", Fn: func(s string) string { return FgLines(s, Green) }},
	"Blue":        {Name: "Blue", Fn: func(s string) string { return FgLines(s, Blue) }},
	"OnSoftYellow": {
		Name: "OnSoftYellow",
		Fn:   func(s string) string { return FgBgLines(s, RGB{0x43, 0x62, 0x12}, SoftYellowBG) },
	},
	"OnSoftGreen": {
		Name: "OnSoftGreen",
		Fn:   func(s string) string { return FgBgLines(s, RGB{0x16, 0x65, 0x34}, SoftGreenBG) },
	},
	"Dim": {Name: "Dim", Fn: func(s string) string { return FgLines(s, DimGray) }},
	"NoColor": {Name: "", Fn: nil},
}

// Convenience aliases you can import in call sites.
var (
	RedText      = Colorizers["Red"]
	GreenText    = Colorizers["Green"]
	BlueText     = Colorizers["Blue"]
	OnSoftYellow = Colorizers["OnSoftYellow"]
	OnSoftGreen  = Colorizers["OnSoftGreen"]
	DimText      = Colorizers["Dim"]
	NoColor      = Colorizers["NoColor"]
)

// for dynamic additions at runtime.
func RegisterColorizer(name string, fn func(string) string) Colorizer {
	c := Colorizer{Name: name, Fn: fn}
	Colorizers[name] = c
	return c
}
