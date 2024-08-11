package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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
		char := rune(fileContentString[i]) // all characters are rune
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
			stringOpen := true
			i++
			String := ""
			for i < len(fileContentString) { // till the end of the string
				if fileContentString[i] == byte(STRING) && fileContentString[i] != '\n' { // if " is found, close string
					stringOpen = false
					break
				}
				if fileContentString[i] == '\n' {
					line++
				}
				String += string(fileContentString[i]) // add all characters to the string
				i++
			}

			if !stringOpen {
				fmt.Println("STRING \""+String+"\"", String)
			} else {
				error = true
				fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.", line)
			}

		case '\n': // for new lines
			line++

		default:
			if isDigit(char) {
				output := ""
				for i+1 < len(fileContentString) && isDigit(rune(fileContentString[i+1])) { // if the next char is digit, add and move ahead
					output += string(fileContentString[i])
					i++
				}

				if fileContentString[i] == '\n' {
					line++
				}

				if i+1 < len(fileContentString) && fileContentString[i+1] == '.' { // if char next to the current digit is dot
					output += string(fileContentString[i])                                     // add the current digit
					if i+2 < len(fileContentString) && isDigit(rune(fileContentString[i+2])) { // if numbe rnext to dot a digit
						output += string(fileContentString[i+1]) // add the dot
						i += 2                                   // move to the digit next to dot
						for i < len(fileContentString) && isDigit(rune(fileContentString[i])) {
							output += string(fileContentString[i])
							i++
						}
						
						if fileContentString[i] == '\n' {
							line++
						}
					}
				} else if i < len(fileContentString) && isDigit(rune(fileContentString[i])) { // if the current digit is not followed by dot
					output += string(fileContentString[i]) // add the current digit
				}

				outputFloat, err := strconv.ParseFloat(output, 64)

				if err != nil {
					error = true
					fmt.Fprintf(os.Stderr, "[line %d] Error: Invalid number: %s\n", line, output)
					continue
				} else {
					formatedOutput := fmt.Sprintf("%.6f", outputFloat)
					formatedOutput = strings.TrimRight(formatedOutput, "0")

					if formatedOutput[len(formatedOutput)-1] == '.' {
						formatedOutput += "0"
					}
					fmt.Println("NUMBER", output, formatedOutput)
				}

				if i < len(fileContentString) && fileContentString[i] == '.' {
					fmt.Println("DOT . null")
				}

			} else if char == '.' {
				fmt.Println("DOT . null")
				i++
			} else if isAlphabet(char) { // if the char is alphabet
				identifier := ""
				if i < len(fileContentString) && isAlphaNumeric(rune(fileContentString[i])) { // if current char is alphanumeric
					identifier += string(fileContentString[i])
				} else {
					fmt.Println("IDENTIFIER", identifier, "null")
				}
			} else {
				if !unicode.IsSpace(char) {
					error = true
					fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %s\n", line, string(char))
				}
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

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}

func isAlphabet(char rune) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

func isAlphaNumeric(char rune) bool {
	return isAlphabet(char) || isDigit(char)
}
