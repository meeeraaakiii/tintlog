package logr

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func sanitizeArg(a any) any {
	switch v := a.(type) {
	case error:
		return v.Error()
	case fmt.Stringer:
		return v.String()
	case []byte:
		// Store as base64 so it's unambiguous and JSON-safe.
		return map[string]any{"__bytes_b64": base64.StdEncoding.EncodeToString(v)}
	// Add other special cases if you need (e.g., complex128 -> string)
	default:
		// Let json.Marshal handle primitives, maps, slices, structs, time.Time, etc.
		// For clearly non-serializable things (func, chan, uintptr), stringify type.
		// We detect them with a type switch (kept simple):
		typeName := fmt.Sprintf("%T", a)
		if strings.HasPrefix(typeName, "func(") || strings.HasPrefix(typeName, "chan ") {
			return typeName
		}
		return a
	}
}

func sanitizeArgs(args []any) []any {
	out := make([]any, len(args))
	for i, a := range args {
		out[i] = sanitizeArg(a)
	}
	return out
}
