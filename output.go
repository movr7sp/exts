package main

import "fmt"

func printoutput(extsmap map[string][]string, p params) {

	if p.showfiles == false && p.showempty == false && p.extfilter[0] == "" {
		for key, _ := range extsmap {
			if key != "" {
				fmt.Println("." + key)
			}
		}

		return
	}

	if p.showfiles == true && p.extfilter[0] == "" {
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

	if p.showempty {
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

			if filter == "" {
				continue
			}
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
