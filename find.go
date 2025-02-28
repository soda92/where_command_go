package main

import (
	"os"
	"path/filepath"
	"strings"
)

func FindFiles(path string, files []os.DirEntry, command string) []Path {
	var ret []Path
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
			ret = append(ret, p)
		}
	}
	return ret
}

func FindCommand(paths []string, command string) List[Path] {
	var found []Path
	for _, path := range paths {
		fileinfo, err := os.Stat(path)
		if err != nil || !fileinfo.IsDir() {
			continue
		}
		files, err := os.ReadDir(path)
		if err != nil {
			continue
		}

		results := FindFiles(path, files, command)
		found = append(found, results...)
	}
	return found
}
