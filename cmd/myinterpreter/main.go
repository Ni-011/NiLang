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
	EQUAL       rune = '='
	BANG        rune = '!'
	LESS        rune = '<'
	GREATER     rune = '>'
	SLASH       rune = '/'
	STRING      rune = '"'
)

var error bool = false

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

	fileContentString := string(fileContents)

	line := 1

	for i := 0; i < len(fileContentString); i++ { // for each char in the file, characterise each token
		char := rune(fileContentString[i])
		switch char {
		case LEFT_PAREN:
			fmt.Println("LEFT_PAREN ( null")

		case RIGHT_PAREN:
			fmt.Println("RIGHT_PAREN ) null")

		case LEFT_BRACE:
			fmt.Println("LEFT_BRACE { null")

		case RIGHT_BRACE:
			fmt.Println("RIGHT_BRACE } null")

		case STAR:
			fmt.Println("STAR * null")

		case DOT:
			fmt.Println("DOT . null")

		case COMMA:
			fmt.Println("COMMA , null")

		case PLUS:
			fmt.Println("PLUS + null")

		case MINUS:
			fmt.Println("MINUS - null")

		case SEMICOLON:
			fmt.Println("SEMICOLON ; null")

		case EQUAL:
			if i+1 < len(fileContentString) && fileContentString[i+1] == byte(EQUAL) {
				fmt.Println("EQUAL_EQUAL == null")
				i++
			} else {
				fmt.Println("EQUAL = null")
			}

		case BANG:
			if i+1 < len(fileContentString) && fileContentString[i+1] == byte(EQUAL) {
				fmt.Println("BANG_EQUAL != null")
				i++
			} else {
				fmt.Println("BANG ! null")
			}

		case LESS:
			if i+1 < len(fileContentString) && fileContentString[i+1] == byte(EQUAL) {
				fmt.Println("LESS_EQUAL <= null")
				i++
			} else {
				fmt.Println("LESS < null")
			}

		case GREATER:
			if i+1 < len(fileContentString) && fileContentString[i+1] == byte(EQUAL) {
				fmt.Println("GREATER_EQUAL >= null")
				i++
			} else {
				fmt.Println("GREATER > null")
			}

		case SLASH:
			if i+1 < len(fileContentString) && fileContentString[i+1] == byte(SLASH) { // if 2 consecutive slashes, it is comment
				for i < len(fileContentString) && fileContentString[i] != '\n' { // till i is not a new line, skip indices
					i++
				}
				line++
			} else {
				fmt.Println("SLASH / null")
			}

		case STRING:
			i++
			String := ""
			for i < len(fileContentString) && fileContentString[i] != byte(STRING) { // till it hits the next "
				String += string(fileContentString[i]) // add all characters to the string
				i++;
			}

			fmt.Println("STRING \"", String, "\" null")

		case '\n': // for new lines
			line++

		default:
			if !unicode.IsSpace(char) { // if char is not a space
				error = true
				fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %s\n", line, string(char))
			}
		}
	}

	fmt.Println("EOF  null")

	if error {
		os.Exit(65)
	} else {
		os.Exit(0)
	}
}
