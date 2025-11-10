// Read log file and print it to stderr.
// Has an option to choose maximum log level.
// This way we can see messages that we missed during the program run.
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/meeeraaakiii/go-tintlog/color"
	"github.com/meeeraaakiii/go-tintlog/logger"
)

func main() {
	logFile := flag.String("file", "", "File path to read.")
	logLevel := flag.Int("level", 99, "Log level. Only print messages with log level <= this.")
	startTimeStr := flag.String("start", "0000/Jan/01 00:00:00", "Start time in --time-format format. Keep empty to read from the beginning of the file.")
	endTimeStr := flag.String("end", "9999/Dec/31 23:59:59", "End time in --time-format format. Keep empty to read to the end of the file.")
	timeFormat := flag.String("time-format", tl.Cfg.TimeFormat, "Time format to use for --start and --end. Default is the same as default logger package time format.")
	tail := flag.Int("tail", -1, "Number of lines to show with --tail.")
	flag.Parse()

	if *logFile == "" {
		fmt.Println("Need to specify --file")
		os.Exit(1)
	}

	// Parse the provided start and end times
	startTime, err := time.Parse(*timeFormat, *startTimeStr)
	if err != nil {
		fmt.Println("Error parsing start time:", err)
		return
	}

	endTime, err := time.Parse(*timeFormat, *endTimeStr)
	if err != nil {
		fmt.Println("Error parsing end time:", err)
		return
	}

	// Read file with combined logic
	err, errMsg := readLogFile(*logFile, tl.LogLevel(*logLevel), startTime, endTime, *tail)
	if err != nil {
		tl.Log(tl.Info, color.Red, "Err: '%s', errMsg: '%s'", err, errMsg)
	}
}

/*
Read --file.
If --tail is set, collect last N lines.
For each line check if it's between startTime and endTime,
and if its logging level is below or equal to --level.
If conditions are satisfied - print this message using fmt
including all other parts of LogLine.
*/
func readLogFile(logFile string, logLevel tl.LogLevel, startTime, endTime time.Time, tailCount int) (err error, errMsg string) {
	// Open the file
	file, err := os.Open(logFile)
	if err != nil {
		return err, "Unable to open file"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var buffer [][]byte

	// First read all lines
	for scanner.Scan() {
		line := append([]byte(nil), scanner.Bytes()...) // copy
		buffer = append(buffer, line)
	}
	if err := scanner.Err(); err != nil {
		return err, "Scanner error while reading file"
	}

	// Trim to last N lines if --tail is set
	if tailCount >= 0 && len(buffer) > tailCount {
		buffer = buffer[len(buffer)-tailCount:]
	}

	// Process and print each line
	for _, line := range buffer {
		err, errMsg = processLogLine(line, logLevel, startTime, endTime)
		if err != nil {
			return err, errMsg
		}
	}

	return nil, ""
}

func processLogLine(logLineBytes []byte, logLevel tl.LogLevel, startTime, endTime time.Time) (err error, errMsg string) {
	var logLine tl.LogLine
	// Unmarshal the JSON into the struct
	err = json.Unmarshal(logLineBytes, &logLine)
	if err != nil {
		return err, fmt.Sprintf("Unable to json.Unmarshal line: '%s'", string(logLineBytes))
	}

	// first check time
	if !(AfterOrEqual(logLine.Time, startTime) && BeforeOrEqual(logLine.Time, endTime)) {
		// skip the line if it's not within our time range
		return nil, ""
	}
	// then check log level
	if logLine.Level > logLevel {
		// skip the line if log level is above specified
		return nil, ""
	}

	// now print it
	printLogLine(logLine)

	return nil, ""
}

func AfterOrEqual(t, u time.Time) bool {
	return t.After(u) || t.Equal(u)
}

func BeforeOrEqual(t, u time.Time) bool {
	return t.Before(u) || t.Equal(u)
}

// pick a colorizer by name, fallback to NoColor
func pickColorizer(name string) color.Colorizer {
	if c, ok := color.Colorizers[name]; ok && c.Fn != nil {
		return c
	}
	return color.Colorizers["NoColor"]
}

func printLogLine(logLine tl.LogLine) {
	// choose colors
	timeColorizer := tl.Cfg.LogTimeColor           // e.g. "Gray" or "#8899aa"
	logLineColorizer := pickColorizer(logLine.Color) // e.g. "Green", "RedBoldBackground"

	// build fields
	timeStr := timeColorizer.Apply(logLine.Time.Format(tl.Cfg.TimeFormat))
	levelStr := logLineColorizer.Apply(logLine.Level.String())

	tidPart := ""
	if logLine.TID > 0 { // NOTE: field is TID (not TId)
		tidPart = "[" + logLineColorizer.Apply(strconv.Itoa(logLine.TID)) + "]"
	}

	// render message (use Format+Args if provided)
	msg := logLine.Format
	if strings.TrimSpace(msg) == "" {
		msg = "%v"
	}

	// color the arguments
	var coloredArgs []any
	if len(logLine.Args) > 0 {
		for _, arg := range logLine.Args {
			coloredArg := tl.PrettyForStderr(arg)
			coloredArgs = append(coloredArgs, logLineColorizer.Apply(coloredArg))
		}
		msg = fmt.Sprintf(msg, coloredArgs...)
	}

	// final line
	// Example: 2025-11-09T18:19:26-05:00 [ERROR][1] message...
	if tidPart != "" {
		fmt.Printf("%s [%s]%s %s\n", timeStr, levelStr, tidPart, msg)
	} else {
		fmt.Printf("%s [%s] %s\n", timeStr, levelStr, msg)
	}
}
