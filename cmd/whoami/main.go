package main

import (
	"fmt"
	"os"
	"os/user"
	"strings"
)

const (
	programName    = "whoami"
	programVersion = "1.0.0"
	authors        = "lnxwizard"
)

func usage() {
	fmt.Printf("Usage: %s [OPTION]...\n", programName)
	fmt.Println("Print the user name associated with the current effective user ID.")
	fmt.Println(`OPTIONS: 
	--help 		display this help and exit
	--version 	output version information and exit
	`)
}

func main() {
	if len(os.Args) == 1 {
		currentUser, err := user.Current()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(currentUser.Username)
	} else {
		switch os.Args[1] {
		case "--help":
			usage()
		case "--version":
			fmt.Printf("%s %s\n", programName, programVersion)
		default:
			if strings.HasPrefix(os.Args[1], "--") {
				fmt.Printf("%s: invalid option '%s'\nTry 'whoami --help' for more information.\n", programName, os.Args[1])
				break
			}

			fmt.Printf("%s: extra operand '%s'\nTry 'whoami --help' for more information.\n", programName, os.Args[1])
		}
	}
}
