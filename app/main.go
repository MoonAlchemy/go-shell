package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var builtins = map[string]struct{}{
	// We add our builtin commands here
	"echo": {},
	"exit": {},
	"type": {},
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		cmd, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		args := strings.Fields(cmd)
		if len(args) < 1 {
			continue
		}

		switch args[0] {
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))

		case "exit":
			os.Exit(0)

		case "type":
			if len(args) < 2 {
				fmt.Println(os.Stderr, "type: missing argument")
				continue
			}

			name := args[1]

			if _, ok := builtins[name]; ok {
				fmt.Printf("%s is a shell builtin \n", name)

			} else if path, err := exec.LookPath(name); err == nil {
				fmt.Printf("%s is %s \n", name, path)
			} else {
				fmt.Printf("%s: not found\n", name)
			}
		default:
			if err := extcmd(args); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

		}
	}

}

func extcmd(args []string) error {
	path, err := exec.LookPath(args[0])

	if err != nil {
		return fmt.Errorf("%s: command not found", args[0])
	}
	cmd := exec.Command(path, args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
