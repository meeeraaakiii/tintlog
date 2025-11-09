package main

import (
	"logger"
)

func main() {
	// logger.InitializeConfig(nil)
	logger.InitializeConfig(&logger.Config{})
	logger.Log(logger.Color, "\n\n\nRegular color text %s, more regular text\nmore text: %s", "text printed with logger.Color\nstillsame logger.Color", "more color text\nnew line colored")
	logger.Log(logger.RedText, "error: %s\n%s", "something", "went wrong")
	logger.Log(logger.OnSoftYellow, "note: retry\nscheduled: %s", "soft yellow bg\nnewline with bg")
}
