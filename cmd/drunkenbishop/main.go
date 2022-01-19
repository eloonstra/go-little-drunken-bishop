package main

import (
	"fmt"
	"github.com/eloonstra/go-little-drunken-bishop/pkg/drunkenbishop"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		help()
		os.Exit(1)
	}

	width := 17
	height := 9
	title := ""
	parameters := 1
	for index, argument := range os.Args {
		switch argument {
		case "--width":
		case "-w":
			width, _ = strconv.Atoi(os.Args[index+1])
			parameters += 2
			break
		case "--height":
		case "-h":
			height, _ = strconv.Atoi(os.Args[index+1])
			parameters += 2
			break

		case "--title":
		case "-t":
			title = os.Args[index+1]
			parameters += 2
			break
		case "--help":
		case "-?":
			help()
			os.Exit(0)
		}
	}

	text := strings.TrimSpace(strings.Join(os.Args[parameters:], " "))
	fmt.Println(drunkenbishop.GenerateRandomArt(width, height, []byte(text), true, title))
}

func help() {
	fmt.Println("Usage: drunken_bishop [options] <text>")
	fmt.Println("Options:")
	fmt.Println("  -w, --width <width>")
	fmt.Println("  -h, --height <height>")
	fmt.Println("  -t, --title <title>")
	fmt.Println("  -?, --help")
}
