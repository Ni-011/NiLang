package main

import (
	"fmt"
	"os"
	"unicode"
)

const (
	LEFT_PAREN  rune = '('
	RIGHT_PAREN rune = ')'
	LEFT_BRACE  rune = '{'
	RIGHT_BRACE rune = '}'
	STAR        rune = '*'
	DOT         rune = '.'
	COMMA       rune = ','
	PLUS        rune = '+'
	MINUS       rune = '-'
	SEMICOLON   rune = ';'
)

var error bool = false

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

	filename := os.Args[2]                     // second argument is the filename
	fileContents, err := os.ReadFile(filename) // read the file
	if err != nil {                            // if there was an error reading the file
		os.Exit(1)
	}

	fileContentString := string(fileContents)

	line := 1
	errors := []string{}
	tokens := []string{}

	for _, char := range fileContentString { // for each char in the file, characterise each token
		switch char {
		case LEFT_PAREN:
			tokens = append(tokens, "LEFT_PAREN ( null")

		case RIGHT_PAREN:
			tokens = append(tokens, "RIGHT_PAREN ) null")

		case LEFT_BRACE:
			tokens = append(tokens, "LEFT_BRACE { null")

		case RIGHT_BRACE:
			tokens = append(tokens, "RIGHT_BRACE } null")

		case STAR:
			tokens = append(tokens, "STAR * null")

		case DOT:
			tokens = append(tokens, "DOT . null")

		case COMMA:
			tokens = append(tokens, "COMMA , null")

		case PLUS:
			tokens = append(tokens, "PLUS + null")

		case MINUS:
			tokens = append(tokens, "MINUS - null")

		case SEMICOLON:
			tokens = append(tokens, "SEMICOLON ; null")

		case '\n': // for new lines
			line++

		default:
			error = true
			if !unicode.IsSpace(char) { // if char is not a space
				errorMsg := fmt.Sprintf("[Line %d] Error: Unxepected character: %s", line, string(char))
				errors = append(errors, errorMsg)
			}
		}
	}

	for _, errorMsg := range errors {
		fmt.Fprintln(os.Stderr, errorMsg)
	}

	for _, token := range tokens {
		fmt.Println(token)
	}

	fmt.Println("EOF  null");

	if error {
		os.Exit(65);
	} else {
		os.Exit(0);
	}
}
