package parser

import (
	"testing"

	"github.com/monban/lispian/token"
)

var parserTests = []struct {
	input  []token.Token
	output List
	err    error
}{
	// An empty list
	// ()
	{
		input: []token.Token{
			token.Start(),
			token.End(),
		},
		output: List{
			T:     EMPTY,
			Items: []Item{},
		},
		err: nil,
	},

	// A literal list containing a single string
	// ("Hello, world")
	{
		input: []token.Token{
			token.Start(),
			token.String("Hello, world"),
			token.End(),
		},
		output: List{
			T:     LITERAL,
			Items: []Item{String("Hello, world")},
		},
		err: nil,
	},

	// A statement list with a string parameter
	// (print "Hello, world")
	{
		input: []token.Token{
			token.Start(),
			token.Statement("print"),
			token.String("Hello, world"),
			token.End(),
		},
		output: List{
			T: STATEMENT,
			Items: []Item{
				Statement("print"),
				String("Hello, world"),
			},
		},
		err: nil,
	},

	// A statement list with two integer parameters
	// (add 1 1)
	{
		input: []token.Token{
			token.Start(),
			token.Statement("add"),
			token.Int("1"),
			token.Int("1"),
			token.End(),
		},
		output: List{
			T: STATEMENT,
			Items: []Item{
				Statement("add"),
				Int(1),
				Int(1),
			},
		},
		err: nil,
	},

	// A list with a sublist
	// (add (add 1 1) 1)
	{
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
		output: List{
			T: STATEMENT,
			Items: []Item{
				Statement("add"),
				List{
					T: STATEMENT,
					Items: []Item{
						Statement("add"),
						Int(1),
						Int(1),
					},
				},
				Int(1),
			},
		},
		err: nil,
	},
}

func TestParse(t *testing.T) {
	for _, tst := range parserTests {
		t.Run("", func(t *testing.T) {
			expected := tst.output
			output, _, err := Parse(tst.input)
			if err != tst.err {
				t.Error(err)
			}

			if output.Equals(expected) {
				t.Logf("%v != %v", output, tst.output)
			} else {
				t.Errorf("%v != %v", output, tst.output)
			}
		})
	}

}
