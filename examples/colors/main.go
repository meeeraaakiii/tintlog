package main

import (
	"github.com/meeeraaakiii/go-tintlog/color"
	"github.com/meeeraaakiii/go-tintlog/logger"
)

func main() {
	useTid := true
	logr.InitializeConfig(&logr.Config{UseTid: &useTid, LogDir: "./log/"})
	logr.Log(logr.Info, color.Green, "\n\n\nRegular color text %s, more regular text\nmore text: %s", "text printed with logr.Color\nstillsame logr.Color", "more color text\nnew line colored")
	logr.Log(logr.Info, color.Red, "error: %s\n%s", "something", "went wrong")
	logr.Log(logr.Info, color.BlueBold, "Here's config:\n'''\n%s\n'''", logr.Cfg)
	logr.Log(logr.Info, color.RedBackground, "error: %s\n%s", "something", "went wrong")
	logr.Log(logr.Info, color.RedBoldBackground, "error: %s\n%s", "something", "went wrong")
}
