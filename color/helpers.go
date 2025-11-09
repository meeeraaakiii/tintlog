package color

import "strings"

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
