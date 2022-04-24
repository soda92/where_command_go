package main

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: w <name>")
		return
	}
	paths := strings.Split(os.Getenv("PATH"), ";")
	found := 0
	for _, path := range paths {
		if fileinfo, err := os.Stat(path); err == nil {
			if fileinfo.IsDir() {
				files, err := ioutil.ReadDir(path)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				for _, file := range files {
					name := filepath.Base(file.Name())
					ext := filepath.Ext(file.Name())
					base := strings.TrimSuffix(name, ext)
					if strings.EqualFold(base, os.Args[1]) {
						fmt.Print(path)
						fmt.Printf("%c", os.PathSeparator)
						color.Set(color.FgGreen)
						fmt.Print(base)
						color.Unset()
						fmt.Println(ext)
						found += 1
					}
				}
			}
		}
	}
	if found == 0 {
		color.Red("Command not found.\n")
	}
}
