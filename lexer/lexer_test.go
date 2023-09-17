package lexer

import (
	"testing"

	"github.com/monban/lispian/token"
)

var lexerTests = []struct {
	input  string
	output []token.Token
}{
	{
		input:  "()",
		output: []token.Token{token.Start(), token.End()},
	},
	{
		input:  "(\"Hello, world\")",
		output: []token.Token{token.Start(), token.String("Hello, world"), token.End()},
	},
}

func TestLexerTokens(t *testing.T) {
	for _, tst := range lexerTests {
		t.Run("", func(t *testing.T) {
			l := Lexer{}
			l.Write([]byte(tst.input))
			expectEqual(t, len(l.tokens), len(tst.output))
			if token.CompareSlice(l.tokens, tst.output) == true {
				t.Logf("%s == %s", tst.input, l.tokens)
			} else {
				t.Errorf("%s != %s", tst.input, l.tokens)
			}
		})
	}
}

func TestStringParsing(t *testing.T) {
	l := Lexer{}
	l.Write([]byte("\""))
	expectEqual(t, l.state, READSTRING)
	l.Write([]byte("foo"))
	expectEqual(t, l.state, READSTRING)
	expectEqual(t, l.partial.String(), "foo")
	l.Write([]byte("\""))
	expectEqual(t, l.state, ROOT)

	expected := token.String("foo")
	actual := l.tokens[0]
	if token.Compare(actual, expected) {
		t.Logf("%s == %s", actual, expected)
	} else {
		t.Errorf("%s != %s", actual, expected)
	}
}

func expectEqual[S comparable](t *testing.T, a S, b S) {
	t.Helper()
	if a == b {
		t.Logf("%#v == %#v\n", a, b)
	} else {
		t.Errorf("%#v != %#v\n", a, b)
	}
}
