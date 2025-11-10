package main

import (
	"github.com/meeeraaakiii/go-tintlog/color"
	"github.com/meeeraaakiii/go-tintlog/logger"
)

func main() {
	useTid := true
	tl.InitializeConfig(&tl.Config{UseTid: &useTid, LogDir: "./log/"})
	tl.Log(tl.Info, color.Green, "\n\n\nRegular color text %s, more regular text\nmore text: %s", "text printed with tl.Color\nstillsame tl.Color", "more color text\nnew line colored")
	tl.Log(tl.Info, color.Red, "error: %s\n%s", "something", "went wrong")
	tl.Log(tl.Info, color.BlueBold, "Here's config:\n'''\n%s\n'''", tl.Cfg)
	tl.Log(tl.Info, color.RedBackground, "error: %s\n%s", "something", "went wrong")
	tl.Log(tl.Info, color.RedBoldBackground, "error: %s\n%s", "something", "went wrong")
}
