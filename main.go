package main

import (
	"fmt"
	"os"
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

	if len(found) == 0 {
		PrintError()
	}
}
