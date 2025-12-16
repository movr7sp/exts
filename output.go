package main

import "fmt"

func printoutput(extsmap map[string][]string, p params) {

	if p.empty {
		for key, value := range extsmap {
			if key == "" {
				fmt.Println("WITHOUT EXTENSION")
			} else {
				continue
			}

			for _, v := range value {
				fmt.Println("\t" + v)
			}
			return
		}
	}

	if p.list {
		if len(p.extfilter) == 0 {
			for key, value := range extsmap {
				if key == "" {
					fmt.Println("WITHOUT EXTENSION")
				} else {
					fmt.Println(value)
				}

				for _, v := range value {
					fmt.Println("\t" + v)
				}
			}
		} else {
			for _, value := range p.extfilter {
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
		fmt.Println("EXTENSIONS: ")
		for key, _ := range extsmap {
			if key == "" {
				continue
			}
			fmt.Println("\t" + "." + key)
		}
	}
}
