package main

import (
	"logger"
	"logger/color"
)

func main() {
	useTid := true
	logger.InitializeConfig(&logger.Config{UseTid: &useTid, LogDir: "./log/"})
	logger.Log(logger.Info, color.GreenText, "\n\n\nRegular color text %s, more regular text\nmore text: %s", "text printed with logger.Color\nstillsame logger.Color", "more color text\nnew line colored")
	logger.Log(logger.Info, color.RedText, "error: %s\n%s", "something", "went wrong")
	logger.Log(logger.Info, color.BlueBoldText, "Here's config:\n'''\n%s\n'''", logger.Cfg)
	logger.Log(logger.Info, color.RedBackgroundText, "error: %s\n%s", "something", "went wrong")
	logger.Log(logger.Info, color.RedBoldBackgroundText, "error: %s\n%s", "something", "went wrong")
}
