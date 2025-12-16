package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type files struct {
	mu   sync.Mutex
	exts map[string][]string
}

func main() {

	var listfiles bool = false

	if len(os.Args) < 2 {
		fmt.Println("USAGE: exts path [-l]")
	}

	var extslice []string

	if len(os.Args) > 2 && os.Args[2] == "-l" {
		if len(os.Args) > 3 {
			extslice = os.Args[3:]
		}
		listfiles = true
	}

	var f files

	f.exts = make(map[string][]string)

	var wg sync.WaitGroup
	//traverse(os.Args[1], &wg, &f)
	wg.Add(1)
	go traverse(".", &wg, &f)
	wg.Wait()

	printoutput(f.exts, listfiles, extslice)

}

func traverse(dir string, wg *sync.WaitGroup, e *files) {
	defer wg.Done()

	nodes, err := os.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	for _, node := range nodes {

		name := node.Name()

		if dir[len(dir)-1] != '/' {
			dir = dir + "/"
		}

		if node.IsDir() {
			wg.Add(1)
			go traverse(dir+name, wg, e)

		} else {
			ext := getext(name)
			e.mu.Lock()
			e.exts[ext] = append(e.exts[ext], dir+name)
			e.mu.Unlock()
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

func printoutput(extsmap map[string][]string, listfiles bool, extslice []string) {
	if listfiles {
		if len(extslice) == 0 {
			for key, value := range extsmap {
				fmt.Println(key)
				for _, v := range value {
					fmt.Println("\t" + v)
				}
			}
		} else {
			for _, value := range extslice {
				_, ok := extsmap[value]
				if ok {
					fmt.Println(value)
					for _, filename := range extsmap[value] {
						fmt.Println("\t" + filename)
					}
				}
			}
		}

	} else {
		for key, _ := range extsmap {
			fmt.Println(key)
		}
	}
}
