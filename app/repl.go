package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
