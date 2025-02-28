package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func ColorPrint(p Path) {
	name := filepath.Base(p.File)
	ext := filepath.Ext(p.File)
	base := strings.TrimSuffix(name, ext)
	fmt.Print(p.Dir)
	fmt.Printf("%c", os.PathSeparator)
	color.Set(color.FgCyan, color.Bold)
	fmt.Print(base)
	color.Unset()
	fmt.Println(ext)
}

func PrintError() {
	color.Red("Command not found.\n")
}
