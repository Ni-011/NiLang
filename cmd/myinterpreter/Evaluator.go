package main

import (
	"fmt"
	"reflect"
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
			return nil, fmt.Errorf("operant must be number.\n[Line %v]", unaryNode.line);

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
			leftNum, leftIsNum := left.(float64);
			rightNum, rightIsNum := right.(float64);
			
			// If both are numbers, add them
			if leftIsNum && rightIsNum {
				return leftNum + rightNum, nil;
			}
			
			leftStr, leftIsStr := left.(string);
			rightStr, rightIsStr := right.(string);
			
			// If both are strings, concatenate them
			if leftIsStr && rightIsStr {
				return leftStr + rightStr, nil;
			}
			
			// Otherwise, it's an error - can't mix types
			return nil, fmt.Errorf("operands must be two numbers or two strings.\n[Line %v]", BinaryNode.line);

		case "-":
			leftNum, ok := left.(float64);
			if ok {
				rightNum, ok := right.(float64);
				if ok {
					return leftNum - rightNum, nil;
				}
			} else {
				return nil, fmt.Errorf("operands must be two numbers.\n[Line %v]", BinaryNode.line);
			}

		case ">": 
			if leftNum, ok := left.(float64); ok {
				if rightNum, ok := right.(float64); ok {
					return leftNum > rightNum, nil;
				}
			} else {
				return nil, fmt.Errorf("'>' operator requires two numbers");
			}

		case ">=": 
			if leftNum, ok := left.(float64); ok {
				if rightNum, ok := right.(float64); ok {
					return leftNum >= rightNum, nil;
				}
			} else {
				return nil, fmt.Errorf("'>=' operator requires two numbers");
			}

		case "<": 
			if leftNum, ok := left.(float64); ok {
				if rightNum, ok := right.(float64); ok {
					return leftNum < rightNum, nil;
				}
			} else {
				return nil, fmt.Errorf("'<' operator requires two numbers");
			}

		case "<=": 
			if leftNum, ok := left.(float64); ok {
				if rightNum, ok := right.(float64); ok {
					return leftNum <= rightNum, nil;
				}
			} else {
				return nil, fmt.Errorf("'<=' operator requires two numbers");
			}

		case "==":
			// Different types are never equal
			if reflect.TypeOf(left) != reflect.TypeOf(right) {
				return false, nil
			}
			
			// Same types can be compared directly
			return left == right, nil

		case "!=":
			// Different types are always not equal
			if reflect.TypeOf(left) != reflect.TypeOf(right) {
				return true, nil
			}
			
			// Same types can be compared directly
			return left != right, nil
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
		
		// If it was a string literal, return it as a string without trying to parse as number
		if LiteralNode.wasString {
			return nodeStr, nil;
		}
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