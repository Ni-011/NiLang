package main

import (
	"fmt"
	"strconv"
)

func Evaluate (source string) (interface{}, error) {
	ast, err := Parse(source);
	if err != nil {
		return nil, err;
	}

	return EvaluateAST(ast.Root);
}

func EvaluateAST(node ASTNode) (interface{}, error) {
	nodeStr := node.String();

	// check if node is unary
	unaryNode, ok := node.(*UnaryNode);
	if ok {
		// evaluate the expression  of unary
		expr, err := EvaluateAST(unaryNode.expr);
		if err != nil {
			return nil, err;
		}

		switch unaryNode.operator {
		case "-" :
			// check if the expression is a number
			num, ok := expr.(float64);
			if ok {
				return -num, nil;
			}
			return nil, fmt.Errorf("cannot apply '-' to non-number: %v", expr);

		case "!":
			// check if the expression is a boolean
			return !isTrue(expr), nil;

		default:
			return nil, fmt.Errorf("unknown unary operator: %s", unaryNode.operator);
		}
	}

	BinaryNode, ok := node.(*BinaryNode);
	if ok {
		left, err := EvaluateAST(BinaryNode.left);
		if err != nil {
			return nil, err;
		}

		right, err := EvaluateAST(BinaryNode.right);
		if err != nil {
			return nil, err;
		}

		switch BinaryNode.operator {
		case "/":
			// check if left is number
			leftNum, ok := left.(float64);
			if ok {
				rightNum, ok := right.(float64);
				if ok {
					if rightNum == 0 {
						return nil, fmt.Errorf("division by 0");
					}

					return leftNum/rightNum, nil;
				}
			}

		case "*":
			leftNum, ok := left.(float64);
			if ok {
				rightNum, ok := right.(float64);
				if ok {
					return leftNum*rightNum, nil;
				}
			}
		
		case "+":
			leftNum, ok := left.(float64);
			if ok {
				rightNum, ok := right.(float64);
				if ok {
					return leftNum + rightNum, nil;
				}
			}

		case "-":
			leftNum, ok := left.(float64);
			if ok {
				rightNum, ok := right.(float64);
				if ok {
					return leftNum - rightNum, nil;
				}
			}
		}

		return nil, fmt.Errorf("unknown binary operator: %s", BinaryNode.operator);
	}

	// group
	groupNode, ok:= node.(*GroupNode);
	if ok {
		return EvaluateAST(groupNode.expression);
	}

	// literal
	LiteralNode, ok := node.(*LiteralNode);
	if ok {
		nodeStr = LiteralNode.value;
	}
	
	switch nodeStr {
		case "nil":
			return nil, nil;
		case "true":
			return true, nil;
		case "false":
			return false, nil;
			
		
		default:
			num, err := strconv.ParseFloat(nodeStr, 64);
			if err == nil {
				return num, nil;
			}

			return nodeStr, nil;
	}
}

func isTrue (value interface{}) bool {
	if value == nil {
		return false;
	}

	// if boolean, return its value
	if bool, ok := value.(bool); ok {
		return bool;
	}

	// everything else is true by default
	return true;
}