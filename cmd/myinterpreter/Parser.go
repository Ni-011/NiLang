package main

import "fmt"

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

	default:
		return nil, fmt.Errorf("unexpected token: %v", token.Type)
	}
}
