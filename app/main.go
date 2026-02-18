package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type commandMap map[string]func(args ...string)

var supportedCommands commandMap

func replay(reader *bufio.Reader) {
	fmt.Print("$ ")
	command, err := reader.ReadString('\n')
	command = command[:len(command)-1]
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
	args := strings.Split(strings.TrimSpace(command), " ")

	if args[0] == "" {
		return
	}
	executorFunction, ok := supportedCommands[args[0]]
	if !ok {
		fmt.Printf("%s: command not found\n", command)
		return
	}
	executorFunction(args[1:]...)
}

func main() {
	supportedCommands = commandMap{
		"exit": func(args ...string) { os.Exit(0) },
		"echo": func(args ...string) { fmt.Printf("%s\n", strings.Join(args, " ")) },
		"type": func(args ...string) {
			for _, arg := range args {
				_, ok := supportedCommands[arg]
				if !ok {
					fmt.Printf("%s: command not found\n", arg)
					continue
				}
				fmt.Printf("%s is a shell builtin\n", arg)
			}
		},
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		replay(reader)
	}
}
