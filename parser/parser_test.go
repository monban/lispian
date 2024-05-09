package parser

import (
	"testing"

	"github.com/monban/lispian/ast"
	"github.com/monban/lispian/token"
)

type ParserTest = struct {
	name   string
	input  []token.Token
	output ast.Element
	err    error
}

var literalListTests = []ParserTest{
	// ()
	{
		name: "an empty list",
		input: []token.Token{
			token.Start(),
			token.End(),
		},
		output: ast.List{},
		err:    nil,
	},

	// ("Hello, world")
	{
		name: "literal list with single string",
		input: []token.Token{
			token.Start(),
			token.String("Hello, world"),
			token.End(),
		},
		output: ast.List{ast.String("Hello, world")},
		err:    nil,
	},
}

var functionCallTests = []ParserTest{
	// (print "Hello, world")
	{
		name: "call with string prarmeter",
		input: []token.Token{
			token.Start(),
			token.Statement("print"),
			token.String("Hello, world"),
			token.End(),
		},
		output: ast.Call{
			Name: "print",
			Parameters: []ast.Element{
				ast.String("Hello, world"),
			},
		},
		err: nil,
	},

	// (add 1 2)
	{
		name: "call with two integer parameters",
		input: []token.Token{
			token.Start(),
			token.Statement("add"),
			token.Int("1"),
			token.Int("2"),
			token.End(),
		},
		output: ast.Call{
			Name: "add",
			Parameters: []ast.Element{
				ast.Int(1),
				ast.Int(2),
			},
		},
		err: nil,
	},
	// (add (add 1 1) 1)
	{
		name: "list with sublist",
		input: []token.Token{
			token.Start(),
			token.Statement("add"),
			token.Start(),
			token.Statement("add"),
			token.Int("1"),
			token.Int("1"),
			token.End(),
			token.Int("1"),
			token.End(),
		},
		output: ast.Call{
			Name: "add",
			Parameters: []ast.Element{
				ast.Call{
					Name: "add",
					Parameters: []ast.Element{
						ast.Int(1),
						ast.Int(1),
					},
				},
				ast.Int(1),
			},
		},
		err: nil,
	},

	// (if true "foo" "bar")
	{
		name: "simple if statement",
		input: []token.Token{
			token.Start(),
			token.Statement("if"),
			token.True(),
			token.String("foo"),
			token.String("bar"),
			token.End(),
		},
		output: ast.Call{
			Name: "if",
			Parameters: []ast.Element{
				ast.True(),
				ast.String("foo"),
				ast.String("bar"),
			},
		},
		err: nil,
	},
}

var invalidListTests = []ParserTest{
	{
		name: "add with unbalanced parameters",
		input: []token.Token{
			token.Start(),
			token.Statement("if"),
			token.True(),
			token.String("foo"),
			token.String("bar"),
		},
		output: ast.NewVoid(),
		err:    UnbalancedParams{},
	},
}

func TestParse(t *testing.T) {
	t.SkipNow()
	tests := append(literalListTests, functionCallTests...)
	tests = append(tests, invalidListTests...)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expected := test.output
			output, err := Parse(test.input)
			if err != test.err {
				t.Error(err)
			}

			if ast.Equal(output, expected) {
				t.Logf("%v == %v", output, test.output)
			} else {
				t.Errorf("%v != %v", output, test.output)
			}
		})
	}

}

func TestParseList(t *testing.T) {
	tests := append(literalListTests, functionCallTests...)
	tests = append(tests, invalidListTests...)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expected := test.output
			output, i, err := parseList(test.input)
			if err != test.err {
				t.Error(err)
			}
			if len(test.input) == i {
				t.Logf("parser received and processed %d tokens", i)
			} else {
				t.Logf("parser received %d tokens but reports it processed %d", len(test.input), i)
			}

			if ast.Equal(output, expected) {
				t.Logf("%v == %v", output, test.output)
			} else {
				t.Errorf("%v != %v", output, test.output)
			}
		})
	}

}

func TestParseElement(t *testing.T) {
	b := []token.Token{
		token.Start(),
		token.String("hello, world"),
		token.End(),
	}
	t.Log(parseElement(b))

}

func TestParseCall(t *testing.T) {
	for _, test := range functionCallTests {
		t.Run(test.name, func(t *testing.T) {
			expected := test.output
			output, i, err := parseCall(test.input)
			if err != test.err {
				t.Error(err)
			}
			if len(test.input) == i {
				t.Logf("parser received and processed %d tokens", i)
			} else {
				t.Logf("parser received %d tokens but reports it processed %d", len(test.input), i)
			}

			if ast.Equal(output, expected) {
				t.Logf("%v == %v", output, test.output)
			} else {
				t.Errorf("%v != %v", output, test.output)
			}
		})
	}
}
