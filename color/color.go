// color.go
package color

import (
	"fmt"
	"strconv"
)

// Color is a hex string like "#RRGGBB".
type Color string

func (c Color) String() string { return string(c) }

// RGB converts the Color (must be "#RRGGBB") to RGB.
func (c Color) RGB() (RGB, error) {
	hex := string(c)
	if len(hex) != 7 || hex[0] != '#' {
		return RGB{}, fmt.Errorf("want #RRGGBB, got %q", hex)
	}
	r, err := strconv.ParseUint(hex[1:3], 16, 8); if err != nil { return RGB{}, err }
	g, err := strconv.ParseUint(hex[3:5], 16, 8); if err != nil { return RGB{}, err }
	b, err := strconv.ParseUint(hex[5:7], 16, 8); if err != nil { return RGB{}, err }
	return RGB{uint8(r), uint8(g), uint8(b)}, nil
}

// MustRGB panics on invalid color; handy for constants.
func (c Color) MustRGB() RGB {
	rgb, err := c.RGB()
	if err != nil { panic(err) }
	return rgb
}
