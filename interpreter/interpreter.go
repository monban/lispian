package interpreter

import (
	"fmt"

	"github.com/monban/lispian/ast"
)

// Evaluate an Element and return the result
func Eval(e ast.Element) ast.Element {
	if l, ok := e.(ast.List); ok {
		// Walk the list and evaluate any sub-lists
		for i, item := range l {
			if item, ok := item.(ast.List); ok {
				l[i] = Eval(item)
			}
		}
		return l
	}
	if e, ok := e.(ast.Call); ok {
		return evalCall(e)
	}
	return e
}

func evalCall(call ast.Call) ast.Element {
	switch call.Name {
	case "add":
		return evalAdd(call.Parameters)
	case "if":
		return evalIf(call.Parameters)
	default:
		err := fmt.Sprintf("unknown statement: '%s'", call.Name)
		panic(err)
	}
}

func evalAdd(l ast.List) ast.Int {
	var sum int
	for _, e := range l {
		e = Eval(e)
		e, ok := e.(ast.Int)
		if !ok {
			panic("unable to parse int")
		}
		sum += int(e)
	}
	return ast.Int(sum)
}

func evalIf(l ast.List) ast.Element {
	p := Eval(l[0])
	pred, ok := p.(ast.Bool)
	if !ok {
		panicMessage := fmt.Sprintf("if statement received non-boolean predicate %v", l[0])
		panic(panicMessage)
	}
	if pred {
		return Eval(l[1])
	}
	return Eval(l[2])
}
