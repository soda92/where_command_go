package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FindCommand(paths []string, command string) List[Path] {
	var found []Path
	for _, path := range paths {
		fileinfo, err := os.Stat(path)
		if err != nil {
			continue
		}
		if !fileinfo.IsDir() {
			continue
		}
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
				var p Path
				p.Dir = path
				p.File = file.Name()
				p.Basename = base
				p.Suffix = ext
				found = append(found, p)
			}
		}
	}
	return found
}
