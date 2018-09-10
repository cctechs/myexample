package main

import (
	"os"
	"path/filepath"
	"flag"
	"log"
	"./download"
)

func main() {
	flag.Parse()
	log.Println(flag.Args())
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	currentDir += "/files/"
	log.Println("download to dir=", currentDir)
	y := youtube.NewYoutube(true)
	arg := flag.Arg(0)
	y.DecodeURL(arg)
	y.StartDownload(currentDir)
}
