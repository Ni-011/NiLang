package main

import (
	"fmt"
	"os"
)

const (
	LEFT_PAREN rune = '('
	RIGHT_PAREN rune = ')'
	LEFT_BRACE rune = '{'
	RIGHT_BRACE rune = '}'
	STAR rune = '*'
	DOT rune = '.'
	COMMA rune = ','
	PLUS rune = '+'
	MINUS rune = '-'
	SEMICOLON rune = ';'
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
		case LEFT_PAREN:
			fmt.Println("LEFT_PAREN ( null");

		case RIGHT_PAREN:
			fmt.Println("RIGHT_PAREN ) null");

		case LEFT_BRACE:
			fmt.Println("LEFT_BRACE { null");

		case RIGHT_BRACE:
			fmt.Println("RIGHT_BRACE } null");

		case STAR:
			fmt.Println("STAR * null");

		case DOT:
			fmt.Println("DOT . null");

		case COMMA:
			fmt.Println("COMMA , null");

		case PLUS:
			fmt.Println("PLUS + null");

		case MINUS:
			fmt.Println("MINUS - null");

		case SEMICOLON:
			fmt.Println("SEMICOLON ; null");

		default:
			fmt.Fprintf(os.Stderr, "[Line[1] Error: Unxepected character: %s", char);
		}
	}

	fmt.Println("EOF  null");
}
