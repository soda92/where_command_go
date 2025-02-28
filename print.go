package main

import (
	"fmt"
	"github.com/fatih/color"
	"path/filepath"
	"strings"
)

func ColorPrint(file string) {
	name := filepath.Base(file)
	ext := filepath.Ext(file)
	base := strings.TrimSuffix(name, ext)
	fmt.Print(path)
	fmt.Printf("%c", os.PathSeparator)
	color.Set(color.FgGreen)
	fmt.Print(base)
	color.Unset()
	fmt.Println(ext)
}

func PrintError() {
	color.Red("Command not found.\n")
}
