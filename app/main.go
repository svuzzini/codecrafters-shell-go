package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		input, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return
		}
		input = strings.TrimSpace(input)
		if input == "exit" {
			return
		}
		if input != "" {
			fmt.Println(input + ": command not found")
		}
		if err == io.EOF {
			return
		}
	}
}
