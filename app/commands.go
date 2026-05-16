package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func initSupportedCommands() commandMap {
	return commandMap{
		"exit": func(args ...string) { os.Exit(0) },
		"echo": func(args ...string) { fmt.Printf("%s\n", strings.Join(args, " ")) },
		"type": func(args ...string) {
			for _, arg := range args {
				_, ok := supportedCommands[arg]
				if !ok {
					path, err := exec.LookPath(arg)
					if err != nil {
						fmt.Printf("%s: not found\n", arg)
						continue
					}
					fmt.Printf("%s is %s\n", arg, path)
					continue
				}
				fmt.Printf("%s is a shell builtin\n", arg)
			}
		},
	}
}
