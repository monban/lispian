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
		input: "()",
		output: []token.Token{
			token.Start(),
			token.End(),
		},
	},
	{
		input: "(\"Hello, world\")",
		output: []token.Token{
			token.Start(),
			token.String("Hello, world"),
			token.End(),
		},
	},
	{
		input: "(\"秘密\")",
		output: []token.Token{
			token.Start(),
			token.String("秘密"),
			token.End(),
		},
	},
}

func TestLexerTokens(t *testing.T) {
	for _, tst := range lexerTests {
		t.Run("", func(t *testing.T) {
			l := Lexer{}
			l.WriteString(tst.input)
			expectEqual(t, len(l.Tokens()), len(tst.output))
			if token.CompareSlice(l.Tokens(), tst.output) == true {
				t.Logf("%s == %s", tst.output, l.Tokens())
			} else {
				t.Errorf("%s != %s", tst.output, l.Tokens())
			}
		})
	}
}

func TestStringParsing(t *testing.T) {
	l := Lexer{}
	l.WriteRune('"')
	expectEqual(t, l.state, READSTRING)
	l.WriteString("foo")
	expectEqual(t, l.state, READSTRING)
	expectEqual(t, l.partial.String(), "foo")
	l.WriteRune('"')
	expectEqual(t, l.state, ROOT)

	expected := token.String("foo")
	actual := l.Tokens()[0]
	if token.Compare(actual, expected) {
		t.Logf("%s == %s", actual, expected)
	} else {
		t.Errorf("%s != %s", actual, expected)
	}
}

func TestUnicodeWriting(t *testing.T) {
	l := Lexer{}
	input := []byte("\"秘密\"")
	expected := token.String("秘密")
	for _, b := range input {
		l.WriteByte(b)
	}
	requireEqual(t, len(l.Tokens()), 1)
	actual := l.Tokens()[0]
	if token.Compare(actual, expected) {
		t.Logf("%s == %s", actual, expected)
	} else {
		t.Errorf("%s != %s", actual, expected)
	}
}

func expectEqual[S comparable](t *testing.T, a S, b S) {
	t.Helper()
	if a == b {
		t.Logf("%v == %v\n", a, b)
	} else {
		t.Errorf("%#v != %#v\n", a, b)
	}
}

func requireEqual[S comparable](t *testing.T, a S, b S) {
	t.Helper()
	if a == b {
		t.Logf("%v == %v\n", a, b)
	} else {
		t.Fatalf("%#v != %#v\n", a, b)
	}
}
