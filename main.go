package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func printTree(path string, prefix string, depth int) {
	if depth == 0 {
		return
	}

	data, _ := os.ReadDir(path)

	for i, file := range data {
		if i == len(data)-1 {
			fmt.Printf("%s└─── %s\n", prefix, file.Name())
			if file.IsDir() {
				printTree(path+"/"+file.Name(), prefix+"│\t", depth-1)
			}
		} else {
			fmt.Printf("%s├─── %s\n", prefix, file.Name())
			if file.IsDir() {
				printTree(path+"/"+file.Name(), prefix+"│\t", depth-1)
			}
		}
	}
}

func main() {
	var depth int
	flag.IntVar(&depth, "n", 3, "Ограничение глубины дерева")
	flag.Parse()
	var path = "."

	if len(os.Args) >= 4 {
		path = os.Args[3]
	}

	if !strings.HasPrefix(path, "/") {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path = filepath.Join(wd, path)
	}

	printTree(path, "", depth)
}
