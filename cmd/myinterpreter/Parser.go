package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Binary operators
var BinaryOperators = map[TokenType]bool {
	STAR:          true,
	SLASH:         true,
	LESS_EQUAL:    true,
	GREATER_EQUAL: true,
	LESS:          true,
	GREATER:       true,
	EQUAL_EQUAL:   true,
	BANG_EQUAL:     true,
}

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

type GroupNode struct {
	expression ASTNode
}

func (g *GroupNode) String() string {
	return fmt.Sprintf("(group %s)", g.expression.String())
}

type UnaryNode struct {
	operator string
	expr ASTNode
}

func (u *UnaryNode) String() string {
	return fmt.Sprintf("(%s %s)", u.operator, u.expr)
}

type BinaryNode struct {
	left ASTNode
	operator string
	right ASTNode
}

func (b *BinaryNode) String() string {
	return fmt.Sprintf("(%s %s %s)", b.operator, b.left, b.right)
}

// error handling
type ParseError struct {
	line int
	lexeme string
	message string
}

// error format
func (e *ParseError) Error() string {
	return fmt.Sprintf("[Line %d] Error at '%s': %s", e.line, e.lexeme, e.message);
}

// main parse function
func (p *Parser) parseExpression() (ASTNode, error) {
	left, err := p.parseBinary()
	if (err != nil) {
		return nil, fmt.Errorf("failed to parse expression: %v", err)
	}

	// check for + - operators
	for p.current < len(p.tokens) {
		operator := p.tokens[p.current];

		if operator.Type != PLUS && operator.Type != MINUS {
			break;
		}
		// consume the operator, move to right
		p.current++;

		right, err := p.parseBinary();
		if err != nil {
			return nil, fmt.Errorf("failed to parse the right term: %v", err)
		}

		left = &BinaryNode{
			left: left,
			operator: operator.Lexeme,
			right: right,
		}
	}

	return left, nil
}

func (p *Parser) parseBinary() (ASTNode, error) {
	// parse left side
	left, err := p.parsePrimary()
	if err != nil {
		return nil, fmt.Errorf("failed to parse the left term: %v", err)
	}

	// check for the operator
	for p.current < len(p.tokens) {
		operator := p.tokens[p.current]

		// if invalid operator, break
		if !BinaryOperators[operator.Type] {
			break
		}

		// if valid operator
		p.current++

		// parse the right term
		right, err := p.parsePrimary()
		if err != nil {
			return nil, fmt.Errorf("failed to parse the right term: %v", err)
		}

		// Update left with the new binary expression
		expr := &BinaryNode{
			left:     left,
			operator: operator.Lexeme,
			right:    right,
		}

		left = expr;
	}
	return left, nil
}

// parses primary expressions, no logical operations
func (p *Parser) parsePrimary() (ASTNode, error) {
	if p.current >= len(p.tokens) {
		return nil, &ParseError{
			line: p.tokens[p.current-1].Line,
			lexeme: p.tokens[p.current-1].Lexeme,
			message: "Expect Expression.",
		}
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

	case LEFT_PAREN:
		// parse the next expr
		expr, err := p.parseExpression();
		if err != nil {
			return nil, err
		}

		// check if the ) exists
		if p.current >= len(p.tokens) || p.tokens[p.current].Type != RIGHT_PAREN {
			return nil, &ParseError{
				line: p.tokens[p.current-1].Line,
				lexeme: p.tokens[p.current-1].Lexeme,
				message: "Expect expression.",
			}
		}
		p.current++ // consume the ')'

		return &GroupNode{expression: expr}, nil;

	case BANG, MINUS:
		operator := token.Lexeme
		expr, err := p.parsePrimary()
		if err != nil {
			return nil, err
		}
		return &UnaryNode{operator: operator, expr: expr}, nil

	default:
		return nil, fmt.Errorf("unexpected token: %v", token.Type)
	}
}