package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func replay(supportedCommands map[string]func(args ...string)) {
	fmt.Print("$ ")
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	command = command[:len(command)-1]
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}
	args := strings.Split(strings.TrimSpace(command), " ")

	if args[0] == "" {
		return
	}
	excutorFunction, ok := supportedCommands[args[0]]
	if !ok {
		fmt.Printf("%s: command not found\n", command)
		return
	}
	excutorFunction(args[1:]...)
}

func main() {
	supportedCommands := map[string]func(args ...string){
		"exit": func(args ...string) { os.Exit(0) },
		"echo": func(args ...string) { fmt.Printf("%s\n", strings.Join(args, " ")) },
	}
	for {
		replay(supportedCommands)
	}
}
