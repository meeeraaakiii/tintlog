// logger/logger.go
package logger

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

var (
	LoggerOutput      io.Writer = os.Stderr
	LoggerOutputMutex sync.Mutex
)

// soft caps to keep stderr snappy
const (
	maxPrettyBytes  = 4096  // cap pretty JSON and text blobs
	maxHexPreview   = 32    // bytes to hex-preview for non-UTF8 []byte
)

func truncate(s string, max int) string {
	if max <= 0 || len(s) <= max {
		return s
	}
	return s[:max] + "…"
}

func prettyForStderr(a any) string {
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
				if n > maxHexPreview { return "…" }
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

// Log prints time (if TimeFormat != ""), [Level], optional [tid], then the message.
// Prints to stderr only when Cfg.LogLevel >= level, but ALWAYS writes JSONL to file
// if LoggerFilePath != "" (colorless), storing only color NAME + original format/args.
func Log(level LogLevel, colorize Colorizer, format string, args ...any) {
	// ----- colored args for stderr -----
	coloredArgs := make([]any, len(args))
	for i, a := range args {
		pretty := prettyForStderr(a)
		if colorize.Fn != nil {
			coloredArgs[i] = colorize.Fn(pretty)
		} else {
			coloredArgs[i] = pretty
		}
	}

	bodyColored := fmt.Sprintf(format, coloredArgs...)
	if !strings.HasSuffix(bodyColored, "\n") {
		bodyColored += "\n"
	}

	// ----- timestamp/prefix for stderr -----
	ts := ""
	if strings.TrimSpace(Cfg.TimeFormat) != "" {
		raw := time.Now().Format(Cfg.TimeFormat)
		if Cfg.LogTimeColor.Name != "" && Cfg.LogTimeColor.Fn != nil {
			raw = Cfg.LogTimeColor.Fn(raw)
		}
		ts = raw + " "
	}

	levelStr := level.String()
	levelStrColored := levelStr
	if colorize.Fn != nil {
		levelStrColored = colorize.Fn(levelStrColored)
	}

	prefix := "[" + levelStrColored + "] "
	tid := 0
	if Cfg.UseTid != nil && *Cfg.UseTid {
		tid = getTid()
		tidStr := strconv.Itoa(tid)
		if colorize.Fn != nil {
			tidStr = colorize.Fn(tidStr)
		}
		prefix = "[" + levelStrColored + "][" + tidStr + "] "
	}

	// ----- ALWAYS write JSONL: original format + sanitized raw args (no ANSI) -----
	writeLogJSONL(LogLine{
		Time:   time.Now(),
		TID:    tid,
		Level:  level,
		Color:  colorize.Name,
		Format: format,
		Args:   sanitizeArgs(args),
	})

	// ----- Print to stderr gated by level -----
	if Cfg.LogLevel >= level {
		LoggerOutputMutex.Lock()
		_, _ = io.WriteString(LoggerOutput, ts+prefix+bodyColored)
		LoggerOutputMutex.Unlock()
	}
}
