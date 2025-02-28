package main

import "strconv"

func Evaluate (source string) (interface{}, error) {
	ast, err := Parse(source);
	if err != nil {
		return nil, err;
	}

	return EvaluateAST(ast.Root);
}

func EvaluateAST(node ASTNode) (interface{}, error) {
	nodeStr := node.String();
	
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