package integration

import (
	"testing"

	"github.com/monban/lispian/ast"
	"github.com/monban/lispian/interpreter"
	"github.com/monban/lispian/lexer"
	"github.com/monban/lispian/parser"
)

var tests = []struct {
	name     string
	input    string
	expected ast.Element
}{
	{
		name:     "add two numbers",
		input:    "(add 2 3)",
		expected: ast.Int(5),
	},
	{
		name:     "nested add 1",
		input:    "(add (add 2 3) 5)",
		expected: ast.Int(10),
	},
	{
		name:     "nested add 2",
		input:    "(add 5 (add 1 3))",
		expected: ast.Int(9),
	},
}

func TestList(t *testing.T) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input := test.input
			expected := test.expected

			l := lexer.Lexer{}
			l.WriteString(input)
			t.Logf("tokens: %v", l.Tokens())

			program, err := parser.Parse(l.Tokens().Tokens)
			if err != nil {
				t.Fatal(err)
			}

			t.Logf("AST: %v", program)
			output := interpreter.Eval(program)

			if output == expected {
				t.Logf("%v == %v", output, expected)
			} else {
				t.Errorf("%v != %v", output, expected)
			}
		})
	}
}

func TestAddition(t *testing.T) {
	input := "(add 1 2)"
	expected := ast.Int(3)

	l := lexer.Lexer{}
	l.WriteString(input)
	t.Logf("tokens: %v", l.Tokens())

	program, err := parser.Parse(l.Tokens().Tokens)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("AST: %v", program)
	output := interpreter.Eval(program)

	if output == expected {
		t.Logf("%v == %v", output, expected)
	} else {
		t.Errorf("%v != %v", output, expected)
	}

}
