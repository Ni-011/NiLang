package main

import (
	"fmt"
	"os"
	"strconv"
)

type TokenType string

// Token types
const (
	LEFT_PAREN    TokenType = "LEFT_PAREN"
	RIGHT_PAREN   TokenType = "RIGHT_PAREN"
	LEFT_BRACE    TokenType = "LEFT_BRACE"
	RIGHT_BRACE   TokenType = "RIGHT_BRACE"
	COMMA         TokenType = "COMMA"
	DOT           TokenType = "DOT"
	MINUS         TokenType = "MINUS"
	PLUS          TokenType = "PLUS"
	SEMICOLON     TokenType = "SEMICOLON"
	SLASH         TokenType = "SLASH"
	STAR          TokenType = "STAR"
	BANG          TokenType = "BANG"
	BANG_EQUAL    TokenType = "BANG_EQUAL"
	EQUAL         TokenType = "EQUAL"
	EQUAL_EQUAL   TokenType = "EQUAL_EQUAL"
	GREATER       TokenType = "GREATER"
	GREATER_EQUAL TokenType = "GREATER_EQUAL"
	LESS          TokenType = "LESS"
	LESS_EQUAL    TokenType = "LESS_EQUAL"
	IDENTIFIER    TokenType = "IDENTIFIER"
	STRING        TokenType = "STRING"
	NUMBER        TokenType = "NUMBER"
	EOF           TokenType = "EOF"

	// Keyword token types
	AND          TokenType = "AND"
	CLASS        TokenType = "CLASS"
	ELSE         TokenType = "ELSE"
	FALSE        TokenType = "FALSE"
	FOR          TokenType = "FOR"
	FUN          TokenType = "FUN"
	IF           TokenType = "IF"
	NIL          TokenType = "NIL"
	OR           TokenType = "OR"
	PRINT        TokenType = "PRINT"
	RETURN       TokenType = "RETURN"
	SUPER        TokenType = "SUPER"
	THIS         TokenType = "THIS"
	TRUE         TokenType = "TRUE"
	VAR          TokenType = "VAR"
	WHILE        TokenType = "WHILE"
)

// each token
type Token struct {
	Type    TokenType
	Lexeme  string
	Literal interface{}
	Line    int
}

// the Lexer
type Lexer struct {
	Source  string
	Tokens  []Token
	Start   int
	Current int
	Line    int
}

// keywords mapped to their token type
var Keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

// constructor function, returns pointer to the struct
func NewLexer(_source string) *Lexer {
	return &Lexer{
		Source:  _source,
		Tokens:  []Token{},
		Start:   0,
		Current: 0,
		Line:    1,
	}
}

// scan all chars in source and returns array of tokens
func (L *Lexer) Scan() []Token {
	for !L.atEnd() { // if not the end, scan that char
		L.Start = L.Current
		L.ScanToken()
	}

	L.Tokens = append(L.Tokens, Token{ // adding EOF token at end
		Type:    EOF,
		Lexeme:  "",
		Literal: nil,
		Line:    L.Line,
	})

	return L.Tokens
}

// scan the current character
func (L *Lexer) ScanToken() {
	char := L.advance()

	switch char {

	case '(':
		L.addToken(LEFT_PAREN)

	case ')':
		L.addToken(RIGHT_PAREN)

	case '{':
		L.addToken(LEFT_BRACE)

	case '}':
		L.addToken(RIGHT_BRACE)

	case ',':
		L.addToken(SEMICOLON)

	case '*':
		L.addToken(STAR)

	case '.':
		L.addToken(DOT)

	case '-':
		L.addToken(MINUS)

	case '+':
		L.addToken(PLUS)

	case ';':
		L.addToken(SEMICOLON)

	case '!':
		if L.match('=') {
			L.addToken(BANG_EQUAL)
		} else {
			L.addToken(BANG)
		}

	case '=':
		if L.match('=') {
			L.addToken(EQUAL_EQUAL)
		} else {
			L.addToken(EQUAL)
		}

	case '<':
		if L.match('=') {
			L.addToken(LESS_EQUAL)
		} else {
			L.addToken(LESS)
		}

	case '>':
		if L.match('=') {
			L.addToken(GREATER_EQUAL)
		} else {
			L.addToken(GREATER)
		}

	case '/':
		if L.match('/') {
			for L.peek() != '\n' && !L.atEnd() {
				L.advance()
			}
		} else {
			L.addToken(SLASH)
		}

	case ' ', '\r', '\t':
        return

	case '\n':
		L.Line++

	case '"':
		L.string()

	default:
		if isDigit(char) {
			L.number()
		} else if isAlpha(char) {
			L.identifier()
		} else {
			L.error(L.Line, char)
		}
	}
}

func (L *Lexer) error(Line int, char byte) {
	fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", Line, char)
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func isAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

func isAlphaNumeric(char byte) bool {
	return isAlpha(char) || isDigit(char)
}

func (L *Lexer) advance() byte {
    current := L.Current
	L.Current++
	return L.Source[current]
}

func (L *Lexer) peek() byte {
	if L.atEnd() {
		return 0
	}
	return L.Source[L.Current]
}

func (L *Lexer) peekNext() byte {
	if L.Current+1 >= len(L.Source) {
		return 0
	}

	return L.Source[L.Current+1]
}

func (L *Lexer) match(expected byte) bool {
	if L.atEnd() {
		return false
	}

	if L.peek() != expected {
		return false
	}

	L.Current++
	return true
}

func (L *Lexer) atEnd() bool {
	return L.Current >= len(L.Source)
}

func (L *Lexer) addToken(_token TokenType) {
	L.addTokenLiteral(_token, nil)
}

func (L *Lexer) addTokenLiteral(_token TokenType, _Literal interface{}) { // add the token to the tokens array
	Lexeme := L.Source[L.Start:L.Current] // get the lexeme

	L.Tokens = append(L.Tokens, Token{
		Type:    _token,
		Lexeme:  Lexeme,
		Literal: _Literal,
		Line:    L.Line,
	})
}

func (L *Lexer) string() {
	for L.peek() != '"' && !L.atEnd() { // while not end of string move ahead
		if L.peek() == '\n' { // new line if multi line string
			L.Line++
		}

		L.advance()
	}

	if L.atEnd() { // if end without closing ", error
		fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", L.Line)
	}

	L.advance() // move over to the ending "

	stringValue := L.Source[L.Start+1 : L.Current-1] //get the string excluding quotes
	L.addTokenLiteral(STRING, stringValue)
}

func (L *Lexer) identifier() {
	for isAlphaNumeric(L.peek()) { // while not end of identifier move ahead
		L.advance()
	}

	identifier := L.Source[L.Start:L.Current] // at the end, get the identifier

	tokenType, exists := Keywords[identifier] // get the token type and check if it exists

	if !exists { // if not, it's an identifier
		tokenType = IDENTIFIER
	}

	L.addToken(tokenType)
}

func (L *Lexer) number() {
	for isDigit(L.peek()) { // while not end of number move ahead
		L.advance()
	}

	if L.peek() == '.' && isDigit(L.peekNext()) { // if next char is . followed by more digits
		L.advance() // move over the "."

		for isDigit(L.peek()) { // till digits exist, move ahead
			L.advance()
		}
	}

	number := L.Source[L.Start:L.Current]                // get the number
	numberLiteral, err := strconv.ParseFloat(number, 64) // convert the string to float

	if err != nil {
		fmt.Fprintf(os.Stderr, "[line %d] Error: Invalid number\n", L.Line)
		return
	}

	L.addTokenLiteral(NUMBER, numberLiteral)
}
