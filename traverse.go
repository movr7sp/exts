package main

import (
	"log"
	"os"
	"strings"
	"sync"
)

type files struct {
	mu      sync.Mutex
	extsmap map[string][]string
}

func traverse(dir string, wg *sync.WaitGroup, f *files) {
	defer wg.Done()

	nodes, err := os.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodes {
		name := node.Name()
		path := dir + "/" + name

		if node.IsDir() {
			wg.Add(1)
			go traverse(path, wg, f)

		} else {
			ext := getext(name)
			f.mu.Lock()
			f.extsmap[ext] = append(f.extsmap[ext], path)
			f.mu.Unlock()
		}
	}
}

func getext(name string) string {
	arr := strings.Split(name, ".")
	if len(arr) == 1 {
		return ""
	}
	return arr[len(arr)-1]
}
