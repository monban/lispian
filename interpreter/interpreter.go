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

func evalStatement(s parser.List) any {
	if s.Items[0] == parser.Statement("add") {
		return evalAdd(s)
	}
	return nil
}

func evalAdd(s parser.List) int {
	var sum int
	for i := 1; i < len(s.Items); i++ {
		e, ok := s.Items[i].(parser.Int)
		if !ok {
			panic("unable to parse int")
		}
		sum += int(e)
	}
	return sum
}
