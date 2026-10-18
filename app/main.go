package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// main is the main function of the shell program. It reads user input from the
// standard input and prints a message indicating that the command was not found. The program continues to prompt for input until the user types "exit" or an EOF is encountered.
func main() {
	// Create a new reader to read from standard input
	reader := bufio.NewReader(os.Stdin)

	// Loop until the user types "exit"
	for {
		// Prompt the user for input
		fmt.Print("$ ")

		// Read the user's input
		input, err := reader.ReadString('\n')

		// Check if there was an error reading the input
		if err != nil && err != io.EOF {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return
		}

		// Remove the newline character from the input
		input = strings.TrimSpace(input)

		// Check if the user typed "exit" to exit the shell
		if input == "exit" {
			return
		}

		fields := strings.Fields(input)

		// Check if the user typed "echo" and print the arguments
		if len(fields) > 0 && fields[0] == "echo" {
			fmt.Println(strings.Join(fields[1:], " "))
		} else if input != "" {
			// Print a message indicating that the command was not found
			fmt.Println(input + ": command not found")
		}

		// Check if the user typed an EOF
		if err == io.EOF {
			return
		}
	}
}
