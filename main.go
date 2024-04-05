package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir := "C:/temp"
	files, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		f, err := files.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			continue
		}

		var filename = f[0].Name()
		var extension = filepath.Ext(filename)
		if extension == ".txt" {
			e := os.Remove(dir + "/" + filename)
			if e != nil {
				log.Fatal(e)
			}
		}
	}
}
