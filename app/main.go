package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	var builtins = map[string]struct{}{
		"echo": {},
		"type": {},
		"exit": {},
	}
	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		args := strings.Fields(command)
		if len(args) == 0 {
			continue
		}
		switch args[0] {
		case "exit":
			return
		case "type":

			if len(args) < 2 {
				fmt.Println("you did not insert a commend to test")
				continue
			}

			name := args[1]
			if _, ok := builtins[name]; ok {
				fmt.Println(name, "is a shell builtin")
			} else if path, err := exec.LookPath(name); err == nil {
				fmt.Println(name, "is located at", path)
			} else {
				fmt.Println(name, "not found")
			}
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))

		default:
			fmt.Printf("%s: command not found\n", args[0])
		}

	}
}
