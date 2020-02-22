package main

import (
	"./download"
	"flag"
	"log"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	log.Println(flag.Args())
	//os.Args[0] = "https://www.youtube.com/watch?v=PFhgjgU36Gg"
	currentDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	//currentDir, _ := filepath.Abs(filepath.Dir(url))
	currentDir += "/files/"
	log.Println("download to dir=", currentDir)
	y := youtube.NewYoutube(true)
	arg := flag.Arg(0)
	y.DecodeURL(arg)
	y.StartDownload(currentDir)
}
