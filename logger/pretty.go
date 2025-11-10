package logr

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode/utf8"
)

// soft caps to keep stderr snappy
const (
	maxPrettyBytes = 4096 // cap pretty JSON and text blobs
	maxHexPreview  = 32   // bytes to hex-preview for non-UTF8 []byte
)

func truncate(s string, max int) string {
	if max <= 0 || len(s) <= max {
		return s
	}
	return s[:max] + "…"
}

func PrettyForStderr(a any) string {
	// Fast path for common string-ish
	switch v := a.(type) {
	case string:
		return v
	case error:
		return v.Error()
	case *time.Time:
		if v == nil {
			return "<nil *time.Time>"
		}
		return v.Format(time.RFC3339)
	case time.Time:
		return v.Format(time.RFC3339)
	case fmt.Stringer:
		return v.String()
	case []byte:
		if utf8.Valid(v) {
			return truncate(string(v), maxPrettyBytes)
		}
		n := len(v)
		preview := v
		if n > maxHexPreview {
			preview = v[:maxHexPreview]
		}
		return fmt.Sprintf("<%d bytes: %s%s>",
			n, strings.ToUpper(hex.EncodeToString(preview)),
			func() string {
				if n > maxHexPreview {
					return "…"
				}
				return ""
			}(),
		)
	}

	// For everything else, attempt pretty JSON first (nice for structs/maps/slices).
	// Avoid obvious non-serializable kinds to skip the allocation just to fail.
	kind := reflect.Indirect(reflect.ValueOf(a)).Kind()
	switch kind {
	case reflect.Func, reflect.Chan, reflect.UnsafePointer:
		return fmt.Sprintf("%T", a)
	}

	// Try JSON (best effort)
	if b, err := json.MarshalIndent(a, "", "  "); err == nil {
		return truncate(string(b), maxPrettyBytes)
	}

	// Fallback: %+v (includes field names for structs)
	return truncate(fmt.Sprintf("%+v", a), maxPrettyBytes)
}
