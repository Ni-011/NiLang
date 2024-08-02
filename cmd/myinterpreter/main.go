package main

import (
	"fmt"
	"os"
)

const (
	LEFT_Paren rune = '(',
	Right_Paren rune = ')',
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1] // first argument is the command

	if command != "tokenize" { // check if the command is tokenize
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2] // second argument is the filename
	fileContents, err := os.ReadFile(filename) // read the file
	if err != nil { // if there was an error reading the file
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	fileContentString := string(fileContents);	

	for _, char := range fileContentString {
		switch char {
		case LEFT_Paren:
			fmt.println("LEFT_PAREN ( null");

		case Right_Paren:
			fmt.println("RIGHT_PAREN ) null");
		}
	}

	fmt.Println("EOF null");
}
