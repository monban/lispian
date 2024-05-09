package ast

import "testing"

func TestStringer(t *testing.T) {
	l := List{
		String("foo"),
		True(),
		List{
			Int(2),
			Int(5),
		},
		Int(1),
		Int(3),
	}
	t.Logf("\n%v\n", l)
}
