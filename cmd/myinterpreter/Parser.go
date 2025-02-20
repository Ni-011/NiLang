package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Parse(source string) (*AST, error) {
	// Lexical analysis
	lexer := NewLexer(string(source))
	tokens := lexer.Scan()

	// create parser isntance
	Parser := &Parser{
		tokens:  tokens,
		current: 0,
	}

	// parse the expression
	expr, err := Parser.parseExpression()
	if err != nil {
		return nil, err
	}

	// return AST
	return &AST{Root: expr}, nil
}

type Parser struct {
	tokens  []Token
	current int
}

type AST struct {
	Root ASTNode
}

type ASTNode interface {
	String() string
}

// converts AST to string
func (ast *AST) String() string {
	return ast.Root.String()
}

// literals booleans, nil etc
type LiteralNode struct {
	value string
}

func (L *LiteralNode) String() string {
	return L.value
}

func (p *Parser) parseExpression() (ASTNode, error) {
	if p.current >= len(p.tokens) {
		return nil, fmt.Errorf("unexpected end of input")
	}

	token := p.tokens[p.current] // get the current token
	p.current++                  // advance ahead

	switch token.Type {
	case TRUE:
		return &LiteralNode{value: "true"}, nil

	case FALSE:
		return &LiteralNode{value: "false"}, nil

	case NIL:
		return &LiteralNode{value: "nil"}, nil
	
	case NUMBER:
		var value string
		// if token is in float format
		// if float has 1 and only 0 after decimal point, return right there
		if strings.Contains(token.Lexeme, ".") {
			for i := 0; i < len(token.Lexeme); i++ {
				if token.Lexeme[i] == '.' {
					if token.Lexeme[i+1] == '0' && token.Lexeme[len(token.Lexeme)-1] == '0' {
						value = token.Lexeme
						return &LiteralNode{value: value}, nil
					} 
				}
			}

			// if decimal value exists, remove trailing 0s
			value = token.Lexeme
			value = strings.TrimRight(value, "0")
			if strings.HasSuffix(value, ".") {
				value = strings.TrimRight(value, ".")
			}
		} else {
			// convert lexeme to float
			number, err := strconv.ParseFloat(token.Lexeme, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid number: %v", token.Lexeme)
			}
			// format float to 1 decimal place
			numberFormatted := fmt.Sprintf("%.1f", number)
			value = numberFormatted
		}
		return &LiteralNode{value: value}, nil

	case STRING:
		return &LiteralNode{value: token.Lexeme[1:len(token.Lexeme)-1]}, nil

	default:
		return nil, fmt.Errorf("unexpected token: %v", token.Type)
	}
}