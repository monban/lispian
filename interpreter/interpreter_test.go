package interpreter

import (
	"reflect"
	"testing"

	"github.com/monban/lispian/parser"
)

var parserTests = []struct {
	input  parser.List
	output any
	err    error
}{
	// Literal lists just evaluate to themselves
	{
		input: parser.List{
			T: parser.LITERAL,
			Items: []parser.Item{
				parser.String("Hello, world"),
			},
		},
		output: parser.List{
			T: parser.LITERAL,
			Items: []parser.Item{
				parser.String("Hello, world"),
			},
		},
		err: nil,
	},
	{
		input: parser.List{
			T: parser.STATEMENT,
			Items: []parser.Item{
				parser.Statement("add"),
				parser.Int(1),
				parser.Int(2),
			},
		},
		output: parser.Int(3),
		err:    nil,
	},
	{
		input: parser.List{
			T: parser.STATEMENT,
			Items: []parser.Item{
				parser.Statement("add"),
				parser.List{
					T: parser.STATEMENT,
					Items: []parser.Item{
						parser.Statement("add"),
						parser.Int(1),
						parser.Int(1),
					},
				},
				parser.Int(1),
			},
		},
		output: parser.Int(3),
		err:    nil,
	},
	{
		input: parser.List{
			T: parser.STATEMENT,
			Items: []parser.Item{
				parser.Statement("if"),
				parser.True(),
				parser.String("foo"),
				parser.String("bar"),
			},
		},
		output: parser.String("foo"),
		err:    nil,
	},
}

func TestEval(t *testing.T) {
	for _, tst := range parserTests {
		t.Run("", func(t *testing.T) {
			expected := tst.output
			output := Eval(tst.input)
			expectedType := reflect.TypeOf(tst.output)
			receivedType := reflect.TypeOf(output)

			if expectedType != receivedType {
				t.Fatalf("%v != %v", expectedType, receivedType)
			}

			var success bool
			switch output := output.(type) {
			case parser.List:
				success = output.Equals(tst.input)
			case parser.Item:
				success = output == expected
			default:
				t.Fatalf("case not handled")
			}

			if success {
				t.Logf("%v == %v", output, expected)
			} else {
				t.Errorf("%v != %v", output, expected)
			}
		})
	}
}
