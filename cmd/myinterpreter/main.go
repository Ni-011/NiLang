package main

import (
	"fmt"
	"os"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1] // first argument is the command

	if command != "tokenize" { // check if the command is tokenize
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]                     // second argument is the filename
	fileContents, err := os.ReadFile(filename) // read the file
	if err != nil {                            // if there was an error reading the file
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	// Tokenize the file
	scanner := NewLexer(string(fileContents))
	tokens := scanner.Scan()

	for _, token := range tokens {
		fmt.Printf("%s %s %v/n", token.Type, token.Lexeme, token.Literal)
	}

	os.Exit(0)
}
