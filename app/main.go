package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	Reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")

		commend, err := Reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(commend[:len(commend)-1], ": command not found\n")
	}
}
