// logger/logger.go
package logger

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	LoggerOutput      io.Writer = os.Stderr
	LoggerOutputMutex sync.Mutex
)
// Log prints time (if TimeFormat != ""), [Level], optional [tid], then the message.
// Prints to stderr only when Cfg.LogLevel >= level, but ALWAYS writes JSONL to file
// if LoggerFilePath != "" (colorless).
func Log(level LogLevel, colorize Colorizer, format string, args ...any) {
	// ----- prepare args (colorless vs colored) -----
	// colorless (for file):
	bodyColorless := fmt.Sprintf(format, args...)
	if !strings.HasSuffix(bodyColorless, "\n") {
		bodyColorless += "\n"
	}

	// colored (for stderr):
	coloredArgs := make([]any, len(args))
	copy(coloredArgs, args)
	if colorize != nil {
		for i, a := range coloredArgs {
			switch v := a.(type) {
			case string:
				coloredArgs[i] = colorize(v)
			case error:
				coloredArgs[i] = colorize(v.Error())
			case fmt.Stringer:
				coloredArgs[i] = colorize(v.String())
			}
		}
	}
	bodyColored := fmt.Sprintf(format, coloredArgs...)
	if !strings.HasSuffix(bodyColored, "\n") {
		bodyColored += "\n"
	}

	// ----- build timestamp/prefix (shared for stderr only) -----
	ts := ""
	if strings.TrimSpace(Cfg.TimeFormat) != "" {
		raw := time.Now().Format(Cfg.TimeFormat)
		if Cfg.LogTimeColor != nil {
			raw = Cfg.LogTimeColor(raw)
		}
		ts = raw + " "
	}

	levelStr := level.String()
	levelStrColored := levelStr
	if colorize != nil {
		levelStrColored = colorize(levelStrColored)
	}

	prefix := "[" + levelStrColored + "] "
	tid := 0
	if Cfg.UseTid != nil && *Cfg.UseTid {
		tid = getTid()
		tidStr := strconv.Itoa(tid)
		if colorize != nil {
			tidStr = colorize(tidStr)
		}
		prefix = "[" + levelStrColored + "][" + tidStr + "] "
	}

	// ----- ALWAYS write JSONL to file (colorless), regardless of log level -----
	writeLogJSONL(LogLine{
		Time:  time.Now(),
		TID:   tid,
		Level: level,
		Msg:   strings.TrimRight(bodyColorless, "\n"), // store without trailing newline
	})

	// ----- Print to stderr only if threshold allows -----
	if Cfg.LogLevel >= level {
		LoggerOutputMutex.Lock()
		_, _ = io.WriteString(LoggerOutput, ts+prefix+bodyColored)
		LoggerOutputMutex.Unlock()
	}
}
