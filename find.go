package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	
	"github.com/fatih/color"
)

func FindCommand(paths []string, command string) int {
	found := 0
	for _, path := range paths {
		if fileinfo, err := os.Stat(path); err == nil {
			if fileinfo.IsDir() {
				files, err := os.ReadDir(path)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				for _, file := range files {
					name := filepath.Base(file.Name())
					ext := filepath.Ext(file.Name())
					base := strings.TrimSuffix(name, ext)
					if strings.EqualFold(base, command) {
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
	return found
}
