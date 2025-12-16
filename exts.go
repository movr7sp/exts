package main

import (
	"log"
	"os"
	"sync"
)

func main() {

	p, err := parseargs(os.Args)

	if err != nil {
		log.Fatal(err)
	}

	var files files

	files.extsmap = make(map[string][]string)

	var wg sync.WaitGroup
	wg.Add(1)

	go traverse(p.root, &wg, &files)

	wg.Wait()

	printoutput(files.extsmap, p)
}
