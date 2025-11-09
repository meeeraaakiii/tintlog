// logger/logr.go
package logr

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/meeeraaakiii/go-tintlog/color"
)

var (
	LoggerOutput      io.Writer = os.Stderr
	LoggerOutputMutex sync.Mutex
)

// Log prints time (if TimeFormat != ""), [Level], optional [tid], then the message.
// Prints to stderr only when Cfg.LogLevel >= level, but ALWAYS writes JSONL to file
// if LoggerFilePath != "" (colorless), storing only color NAME + original format/args.
func Log(level LogLevel, colorize color.Colorizer, format string, args ...any) {
	// ----- colored args for stderr -----
	coloredArgs := make([]any, len(args))
	for i, a := range args {
		pretty := PrettyForStderr(a)
		coloredArgs[i] = colorize.Apply(pretty)
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
