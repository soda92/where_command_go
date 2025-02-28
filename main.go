package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"runtime"
	"strings"
)

var paths []string

func init() {
	sep := ";"
	if runtime.GOOS == "linux" {
		sep = ":"
	}
	paths = strings.Split(os.Getenv("PATH"), sep)
}

func CheckArgs() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: w <name>")
		os.Exit(-1)
	}
}

func main() {
	CheckArgs()
	found := FindCommand(paths, os.Args[1])

	if found == 0 {
		color.Red("Command not found.\n")
	}
}
