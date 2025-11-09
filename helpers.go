package logger

import (
	"log"
	"runtime"
	"strconv"
	"strings"
)

// get goroutine id
// gotta have this, it will slow things down but we need to track
// in which goroutine thing happens
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
