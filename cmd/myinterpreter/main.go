package main

import (
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2] // takes the file name from cli
	fileContents, err := os.ReadFile(filename) // Reads the file contents into a byte array
	// if error, print the rror and exit
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	
	if len(fileContents) > 0 {
		panic("Scanner not implemented") // if file has some data, scan it, scanner not implemented yet
	} else {
		fmt.Println("EOF  null") // Placeholder, remove this line when implementing the scanner
	}
}
