package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	for _, flag := range os.Args[1:] {
		if flag[0:2] != "-I" {
			continue
		}
		incdir := flag[2:]
		if _, err := os.Stat(incdir); err != nil {
			continue
		}
		_ = filepath.Walk(incdir,
			func(path string, info os.FileInfo, err error) error {
				if info.IsDir() {
					return nil
				}
				header, err := filepath.Rel(incdir, path)
				if err != nil {
					return nil
				}
				fmt.Println(header)
				return nil
			})
	}
}
