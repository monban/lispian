package ast

import "testing"

var equalTests = []struct {
	name   string
	a      Element
	b      Element
	output bool
}{
	{
		name:   "empty list",
		a:      List{},
		b:      List{},
		output: true,
	},
	{
		name: "simple call",
		a: Call{
			Name: "add",
			Parameters: []Element{
				Int(1),
				Int(2),
			},
		},
		b: Call{
			Name: "add",
			Parameters: []Element{
				Int(1),
				Int(2),
			},
		},
		output: true,
	},
	{
		name: "nested list",
		a: List{
			List{
				Int(1),
				Int(42),
			},
			String("hello"),
			NewVoid(),
		},
		b: List{
			List{
				Int(1),
				Int(42),
			},
			String("hello"),
			NewVoid(),
		},
		output: true,
	},
}

func TestEqual(t *testing.T) {
	for _, test := range equalTests {
		t.Run(test.name, func(t *testing.T) {
			if Equal(test.a, test.b) {
				t.Logf("%v == %v", test.a, test.b)
			} else {
				t.Errorf("%v != %v", test.a, test.b)
			}
		})
	}

}
