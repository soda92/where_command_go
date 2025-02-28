package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindCommand(paths []string, command string) []string {
	var found []string
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
						found = append(found, file.Name())
					}
				}
			}
		}
	}
	return found
}
