// color.go
package color

import (
	"fmt"
)

// Color is a hex string like "#RRGGBB".
type Color string

func (c Color) String() string { return string(c) }

// MustRGB panics on invalid color; handy for constants.
func (c Color) MustRGB() RGB {
	rgb, err := c.RGB()
	if err != nil { panic(err) }
	return rgb
}

/*
RGB parses the Color hex string and returns its 24-bit components.
Accepts "#RRGGBB", "#RRGGBBAA", "#RGB", or "#RGBA" (alpha, if present, is ignored).
Hex digits are case-insensitive. Returns an error for malformed input.
*/
func (c Color) RGB() (RGB, error) {
	s := string(c)
	if len(s) < 4 || s[0] != '#' {
		return RGB{}, fmt.Errorf("want #RRGGBB, #RRGGBBAA, #RGB or #RGBA, got %q", s)
	}

	switch len(s) {
	case 7: // #RRGGBB
		r, ok := hx2(s[1], s[2]); if !ok { return RGB{}, fmt.Errorf("bad R in %q", s) }
		g, ok := hx2(s[3], s[4]); if !ok { return RGB{}, fmt.Errorf("bad G in %q", s) }
		b, ok := hx2(s[5], s[6]); if !ok { return RGB{}, fmt.Errorf("bad B in %q", s) }
		return RGB{r, g, b}, nil

	case 9: // #RRGGBBAA (ignore alpha)
		r, ok := hx2(s[1], s[2]); if !ok { return RGB{}, fmt.Errorf("bad R in %q", s) }
		g, ok := hx2(s[3], s[4]); if !ok { return RGB{}, fmt.Errorf("bad G in %q", s) }
		b, ok := hx2(s[5], s[6]); if !ok { return RGB{}, fmt.Errorf("bad B in %q", s) }
		// aa, _ := hx2(s[7], s[8]) // ignored
		return RGB{r, g, b}, nil

	case 4: // #RGB
		r, ok := hx1(s[1]); if !ok { return RGB{}, fmt.Errorf("bad R in %q", s) }
		g, ok := hx1(s[2]); if !ok { return RGB{}, fmt.Errorf("bad G in %q", s) }
		b, ok := hx1(s[3]); if !ok { return RGB{}, fmt.Errorf("bad B in %q", s) }
		return RGB{r, g, b}, nil

	case 5: // #RGBA (ignore alpha)
		r, ok := hx1(s[1]); if !ok { return RGB{}, fmt.Errorf("bad R in %q", s) }
		g, ok := hx1(s[2]); if !ok { return RGB{}, fmt.Errorf("bad G in %q", s) }
		b, ok := hx1(s[3]); if !ok { return RGB{}, fmt.Errorf("bad B in %q", s) }
		// a := s[4] // ignored
		return RGB{r, g, b}, nil
	}

	return RGB{}, fmt.Errorf("unsupported color format %q", s)
}

// parse one hex digit, return doubled byte (e.g., 'a' -> 0xaa)
func hx1(c byte) (byte, bool) {
	n, ok := nib(c); if !ok { return 0, false }
	return n<<4 | n, true
}

// parse two hex digits into a byte
func hx2(a, b byte) (byte, bool) {
	hi, ok1 := nib(a); if !ok1 { return 0, false }
	lo, ok2 := nib(b); if !ok2 { return 0, false }
	return (hi<<4 | lo), true
}

func nib(c byte) (byte, bool) {
	switch {
	case '0' <= c && c <= '9': return c - '0', true
	case 'a' <= c && c <= 'f': return c - 'a' + 10, true
	case 'A' <= c && c <= 'F': return c - 'A' + 10, true
	default: return 0, false
	}
}
