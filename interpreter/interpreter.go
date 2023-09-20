package interpreter

import (
	"github.com/monban/lispian/parser"
)

type ReturnType interface {
	parser.Null | parser.List | parser.Int | parser.String
}

// Evaluate a List and return the result Item
func Eval(l parser.List) any {
	if l.T == parser.LITERAL {
		return l
	}
	if l.T == parser.STATEMENT {
		return evalStatement(l)
	}
	return parser.Null{}
}

func evalStatement(l parser.List) any {
	switch l.Items[0] {
	case parser.Statement("add"):
		return evalAdd(l)
	default:
		panic("unknown statement")
	}
}

func evalAdd(l parser.List) int {
	var sum int
	for i := 1; i < len(l.Items); i++ {
		// TODO: if this is a list, we should check if it evaluates to an Int
		e, ok := l.Items[i].(parser.Int)
		if !ok {
			panic("unable to parse int")
		}
		sum += int(e)
	}
	return sum
}
