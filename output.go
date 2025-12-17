package main

import "fmt"

func printoutput(extsmap map[string][]string, p params) {

	if p.list == false && p.empty == false && p.extfilter[0] == "" {
		flag := 0
		for key, _ := range extsmap {
			if key != "" {
				fmt.Println("." + key)
			} else {
				flag = 1
			}
		}

		if flag == 1 {
			fmt.Println("----")
			fmt.Println("some files without extension")
		}
		return
	}

	if p.list && p.extfilter[0] == "" {
		for key, value := range extsmap {
			if key == "" {
				fmt.Println("Without extension: ")
			} else {
				fmt.Println("." + key)
			}

			for _, v := range value {
				fmt.Println("\t" + v)
			}
		}
		return
	}

	if p.empty {
		for key, value := range extsmap {
			if key == "" {
				fmt.Println("Without extension: ")
			} else {
				continue
			}

			for _, v := range value {
				fmt.Println("\t" + v)
			}
		}
		return
	}

	if p.extfilter[0] != "" {
		for _, filter := range p.extfilter {

			values, ok := extsmap[filter]

			if ok {
				fmt.Println("." + filter)

				for _, v := range values {
					fmt.Println("\t" + v)
				}

			}

		}
		return
	}

}
