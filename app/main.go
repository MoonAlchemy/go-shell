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
	"pwd":  {},
	"cd":   {},
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		input = strings.TrimRight(input, "\n")
		args := parseArgs(input)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))

		case "cd":
			if len(args) < 2 {
				fmt.Fprintln(os.Stderr, "cd: missing argument")
				continue
			}
			dir := args[1]
			if dir == "~" {
				dir = os.Getenv("HOME")
			}
			if err := os.Chdir(dir); err != nil {
				fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", dir)
			}
		case "pwd":
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			fmt.Printf("%s\n", pwd)
		case "exit":
			os.Exit(0)

		case "type":
			if len(args) < 2 {
				fmt.Fprintln(os.Stderr, "type: missing argument")
				continue
			}

			name := args[1]

			if _, ok := builtins[name]; ok {
				fmt.Printf("%s is a shell builtin\n", name)

			} else if path, err := exec.LookPath(name); err == nil {
				fmt.Printf("%s is %s\n", name, path)
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
	_, err := exec.LookPath(args[0])

	if err != nil {
		return fmt.Errorf("%s: command not found", args[0])
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func parseArgs(input string) []string {

	const (
		normalmode = 0
		singlemode = 1
		doublemode = 2
	)
	var mode int
	var escaped bool
	var current strings.Builder
	var args []string

	for _, ch := range input {
		if escaped {
			current.WriteRune(ch)
			escaped = false
			continue
		} else if ch == '\\' && mode == normalmode {
			escaped = true
			continue
		} else if ch == '\'' && mode == normalmode {
			mode = singlemode
		} else if ch == '\'' && mode == singlemode {
			mode = normalmode
		} else if ch == '"' && mode == normalmode {
			mode = doublemode
		} else if ch == '"' && mode == doublemode {
			mode = normalmode
		} else if ch == ' ' && mode == normalmode {
			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
		} else {
			current.WriteRune(ch)
		}
	}
	if current.Len() > 0 {
		args = append(args, current.String())
	}
	return args
}
