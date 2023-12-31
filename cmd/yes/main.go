package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	programName    = "yes"
	programVersion = "1.0.0"
	authors        = "lnxwizard"
)

func usage() {
	fmt.Printf(`Usage: %s [STRING]...  
  or: %s OPTION
Repeatedly output a line with all specified STRING(s), or 'y'.

OPTIONS:
	--help		display this help and exit
	--version	output version information and exit
`, programName, programName)
}

func main() {
	if len(os.Args) == 1 {
		for {
			fmt.Println("y")
		}
	} else {
		if os.Args[1] == "--help" {
			usage()
		} else if os.Args[1] == "--version" {
			fmt.Printf("%s (coreutils) %s\n", programName, programVersion)
		} else {
			for {
				argsWithoutProg := strings.Join(os.Args[1:], " ")
				fmt.Println(argsWithoutProg)
			}
		}
	}
}
