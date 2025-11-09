package logger

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// get goroutine id
func getTid() (tid int) {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	tid, err := strconv.Atoi(idField)
	if err != nil {
		log.Printf("Cannot get goroutine id, Err: %s", err.Error())
		return -1
	}
	return tid
}

type LogLine struct {
	Time   time.Time   `json:"time"`
	TID    int         `json:"tid,omitempty"`
	Level  LogLevel    `json:"level"`
	Color  string      `json:"color,omitempty"` // e.g., "Green"
	Format string      `json:"format"`          // original format string
	Args   []any       `json:"args"`            // sanitized args (no ANSI)
}

// --- file I/O ---

func ensureLogFileOpen() error {
	LoggerFileMutex.Lock()
	defer LoggerFileMutex.Unlock()

	if LoggerFilePath == "" {
		return nil
	}
	if LoggerFile != nil {
		return nil
	}
	f, err := os.OpenFile(LoggerFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	LoggerFile = f
	return nil
}

func writeLogJSONL(line LogLine) {
	if LoggerFilePath == "" {
		return
	}
	if err := ensureLogFileOpen(); err != nil {
		// best-effort
		return
	}
	b, err := json.Marshal(line)
	if err != nil {
		return
	}
	LoggerFileMutex.Lock()
	_, _ = LoggerFile.Write(append(b, '\n'))
	LoggerFileMutex.Unlock()
}

// --- arg sanitization for JSON ---

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
