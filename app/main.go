package main

import (
	"bufio"
	"os"
)

type commandMap map[string]func(args ...string)

var supportedCommands commandMap

func main() {
	pathenv := os.Getenv("PATH")
	if pathenv == "" {
		pathenv = "/bin:/usr/bin:/usr/local/bin"
	}
	supportedCommands = initSupportedCommands()
	reader := bufio.NewReader(os.Stdin)
	for {
		replay(reader)
	}
}
