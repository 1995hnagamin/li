package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func list(prefix string, dirpath string) error {
	dir, err := os.Open(dirpath)
	if err != nil {
		return err
	}
	children, err := dir.Readdir(0)
	if err != nil {
		return err
	}
	for _, child := range children {
		name := child.Name()
		if child.IsDir() {
			err := list(prefix+name+"/", dirpath+"/"+name)
			if err != nil {
				return err
			}
		} else {
			fmt.Println(prefix + name)
		}
	}
	return nil
}

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
