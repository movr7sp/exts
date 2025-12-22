package main

import (
	"errors"
	"fmt"
	"os"
)

type params struct {
	root      string
	showfiles bool
	showempty bool
	extfilter []string
}

func parseargs(args []string) (params, error) {

	var p params
	p.root = ""
	p.extfilter = make([]string, 10)

	j := 0

	help := `
USAGE: exts [PATH] [-l] [-e] [EXTENSIONS]
OPTIONS:
		-l
			list all files grouped by extension
		-e
			list only files without extensions`

	for i := 1; i < len(args); i++ {

		arg := args[i]

		switch arg {
		case "-h", "--help":
			{
				fmt.Println(help)
				os.Exit(1)
			}
		case "-l":
			{
				p.showfiles = true
			}
		case "-e":
			{
				p.showempty = true
			}
		default:
			{
				if arg[0] == '-' {
					return params{}, errors.New("Unknown argmuent: " + arg)
				}

				if i == 1 {
					_, err := os.ReadDir(arg)
					if err == nil {
						p.root = arg
					} else {
						p.root = "."
						p.extfilter[j] = arg
						j++
					}

				} else {
					p.extfilter[j] = arg
					j++
				}
			}
		}
	}

	if p.root == "" {
		p.root = "."
	}

	return p, nil

}
