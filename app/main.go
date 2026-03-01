package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	Reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")

		commend, err := Reader.ReadString('\n')
		commend = strings.TrimSpace(commend)
		if commend == "exit" {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(commend, ": command not found\n")
	}
}
