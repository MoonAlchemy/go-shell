package maim

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
)

func main() {
	
	reader := bufio.NewReadWriter(os.Stdin)
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
			break
		case "type":
			name := args[1]
			if len(args) < 2{
				fmt.Println("you did not insert a commend to test")
			if _, ok := builtins[name], ok {
				fmt.Println(name, "is a shell builtin command")
			} else {
				fmt.Println(name, "is not a shell builtin")
			}
		case "echo":
			fmt.Println(args[1:])
			}
		
		}



		
	}