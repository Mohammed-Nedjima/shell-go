package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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
		_, err := exec.LookPath(args[0])
		if err != nil {
			fmt.Printf("%s: command not found\n", args[0])
			return
		}
		cmd := exec.Command(args[0], args[1:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		err = cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error executing command:", err)
		}
	} else {
		executorFunction(args[1:]...)
	}
}
