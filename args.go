package main

import (
	"errors"
	"fmt"
	"os"
)

type params struct {
	root      string
	list      bool
	empty     bool
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
				p.list = true
			}
		case "-e":
			{
				p.empty = true
			}
		default:
			{
				if arg[0] == '-' {
					return params{}, errors.New("Unknown argmuent: " + arg)
				}

				if i == 1 {
					p.root = arg
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
