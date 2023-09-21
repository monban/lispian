package interpreter

import (
	"fmt"

	"github.com/monban/lispian/parser"
)

type ReturnType interface {
	parser.Null | parser.List | parser.Int | parser.String
}

// Evaluate a List and return the result Item
func Eval(l parser.List) parser.Item {
	// Walk the list and evaluate any sub-lists
	for i, item := range l.Items {
		if item, ok := item.(parser.List); ok {
			l.Items[i] = Eval(item)
		}

	}
	if l.T == parser.LITERAL {
		return l
	}
	if l.T == parser.STATEMENT {
		return evalStatement(l)
	}
	return parser.Null{}
}

func evalStatement(l parser.List) parser.Item {
	switch l.Items[0] {
	case parser.Statement("add"):
		return evalAdd(l)
	default:
		err := fmt.Sprintf("unknown statement: '%#v'", l.Items[0])
		panic(err)
	}
}

func evalAdd(l parser.List) parser.Int {
	var sum int
	for i := 1; i < len(l.Items); i++ {
		// TODO: if this is a list, we should check if it evaluates to an Int
		e, ok := l.Items[i].(parser.Int)
		if !ok {
			panic("unable to parse int")
		}
		sum += int(e)
	}
	return parser.Int(sum)
}
