package main

import (
	"fmt"
	"os"
)

func main() {
	// check if at least 3 commands
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	//check if command is tokenize otherwise error exit1
	command := os.Args[1]

	if command == "tokenize" {
		// get file form args and read it, if error exit1
		filename := os.Args[2]
		fileContents, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}

		// tokenize the file using scanner, get all tokens
		scanner := NewLexer(string(fileContents))
		tokens := scanner.Scan()

		// Print all tokens first
		for _, token := range tokens {
			fmt.Fprintf(os.Stdout, "%s %s %s\n", token.Type, token.Lexeme, token.LiteralString())
		}

		// Then check for errors and exit if needed
		if scanner.hadError {
			os.Exit(65)
		}
	} else if command == "parse" {
		filename := os.Args[2]
		fileContents, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}

		ast, err := Parse(string(fileContents))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(65)
		}
		fmt.Println(ast.Root.String())
	} else if command == "evaluate" {
		filename := os.Args[2];
		fileContents, err := os.ReadFile(filename);
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1);
		}

		result, err := Evaluate(string(fileContents))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error evaluating file: %v\n", err)
			os.Exit(70);
		}

		if result == nil {
			fmt.Println("nil")
		} else {
			fmt.Println(result)
		}
	} else {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	os.Exit(0)
}
