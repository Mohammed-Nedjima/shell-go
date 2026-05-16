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
		"pwd": func(args ...string) {
			wd, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error getting current working directory:", err)
				return
			}
			fmt.Println(wd)
		},
		"cd": func(args ...string) {
			// case of the path not existing
			_, err := os.Stat(args[0])
			if err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", args[0])
				return
			}
			err = os.Chdir(args[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error changing working directory:", err)
			}
		},
	}
}
