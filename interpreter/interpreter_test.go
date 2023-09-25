package interpreter

import (
	"reflect"
	"testing"

	"github.com/monban/lispian/ast"
)

var parserTests = []struct {
	name   string
	input  ast.Element
	output ast.Element
	err    error
}{
	{
		name: "literal list",
		input: ast.List{
			ast.String("Hello, world"),
		},
		output: ast.List{
			ast.String("Hello, world"),
		},
		err: nil,
	},
	{
		name: "add two numbers",
		input: ast.Call{
			Name: "add",
			Parameters: ast.List{
				ast.Int(1),
				ast.Int(2),
			},
		},
		output: ast.Int(3),
		err:    nil,
	},
	{
		name: "nested add",
		input: ast.Call{
			Name: "add",
			Parameters: ast.List{
				ast.Call{
					Name: "add",
					Parameters: ast.List{
						ast.Int(1),
						ast.Int(1),
					},
				},
				ast.Int(1),
			},
		},
		output: ast.Int(3),
		err:    nil,
	},
	{
		name: "if",
		input: ast.Call{
			Name: "if",
			Parameters: ast.List{
				ast.True(),
				ast.String("foo"),
				ast.String("bar"),
			},
		},
		output: ast.String("foo"),
		err:    nil,
	},
}

func TestEval(t *testing.T) {
	for _, tst := range parserTests {
		t.Run(tst.name, func(t *testing.T) {
			expected := tst.output
			output := Eval(tst.input)
			expectedType := reflect.TypeOf(tst.output)
			receivedType := reflect.TypeOf(output)

			if expectedType != receivedType {
				t.Fatalf("%v != %v", receivedType, expectedType)
			}

			var success bool
			switch output := output.(type) {
			case ast.List:
				success = ast.Equal(output, tst.input)
			case ast.Element:
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
