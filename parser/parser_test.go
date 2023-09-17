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
}

func TestParse(t *testing.T) {
	for _, tst := range parserTests {
		t.Run("", func(t *testing.T) {
			expected := tst.output
			output, err := Parse(tst.input)
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
