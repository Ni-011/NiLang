package main

import "fmt"

func Evaluate (source string) (interface{}, error) {
	ast, err := Parse(source);
	if err != nil {
		return nil, err;
	}

	return EvaluateAST(ast.Root);
}

func EvaluateAST(node ASTNode) (interface{}, error) {
	n := node.String();
	
	switch n {
		case "nil":
			return nil, nil;
		case "true":
			return true, nil;
		case "false":
			return false, nil;
		
		default:
			return nil, fmt.Errorf("unknown node: %s", n);
	}
}