package main

import (
	"logger"
)

func main() {
	logger.Log(logger.Color, "Regular color text %s, more regular text\nmore text: %s", "text printed with logger.Color\nstillsame logger.Color", "more color text\nnew line colored")
	logger.Log(logger.RedText, "\n\nerror: %s\n%s", "something", "went wrong")
	logger.Log(logger.OnSoftYellow, "note: retry\nscheduled")
}
