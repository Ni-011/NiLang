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

			leftStr, leftIsStr := left.(string);
			rightStr, rightIsStr := right.(string);

			// if either one is string, convert other and concatinate as strings
			if leftIsStr || rightIsStr {
				if !leftIsStr {
					leftStr = fmt.Sprintf("%v", left);
				}

				if !rightIsStr {
					rightStr = fmt.Sprintf("%v", right);
				}

				return leftStr + rightStr, nil;
			}


		case "-":
			leftNum, ok := left.(float64);
			if ok {
				rightNum, ok := right.(float64);
				if ok {
					return leftNum - rightNum, nil;
				}
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
			// Check if one is a number and one is a string
			leftNum, leftIsNum := left.(float64)
			rightNum, rightIsNum := right.(float64)
			leftStr, leftIsStr := left.(string)
			rightStr, rightIsStr := right.(string)
			
			// If one is a number and one is a string, they're not equal
			if (leftIsNum && rightIsStr) || (leftIsStr && rightIsNum) {
				return false, nil
			}
			
			// If both are numbers, compare them
			if leftIsNum && rightIsNum {
				return leftNum == rightNum, nil
			}
			
			// If both are strings, compare them
			if leftIsStr && rightIsStr {
				return leftStr == rightStr, nil
			}
			
			// For other types, use direct comparison
			return left == right, nil

		case "!=":
			if reflect.TypeOf(left) != reflect.TypeOf(right) {
				return true, nil;
			}

			return left != right, nil;
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

func isSameType (left, right interface{}) bool {
	// check if number or string
	_, leftIsNum := left.(float64)
    _, rightIsNum := right.(float64)
    _, leftIsStr := left.(string)
    _, rightIsStr := right.(string)

	if leftIsNum && rightIsNum {
		return true;
	}

	if leftIsStr && rightIsStr {
		return true;
	}

	if (leftIsNum && rightIsStr) || (leftIsStr && rightIsNum) {
		return false;
	}

	return reflect.TypeOf(left) == reflect.TypeOf(right);
}