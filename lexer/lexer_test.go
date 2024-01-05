package lexer

import (
	"slices"
	"testing"

	"github.com/monban/lispian/token"
)

var lexerTests = []struct {
	input  string
	output *token.List
}{
	{
		input: "()",
		output: token.NewList([]token.Token{
			token.Start(),
			token.End(),
		}),
	},
	{
		input: "(\"Hello, world\")",
		output: token.NewList([]token.Token{
			token.Start(),
			token.String("Hello, world"),
			token.End(),
		}),
	},
	{
		input: "(\"秘密\")",
		output: token.NewList([]token.Token{
			token.Start(),
			token.String("秘密"),
			token.End(),
		}),
	},
	{
		input: "(5)",
		output: token.NewList([]token.Token{
			token.Start(),
			token.Int("5"),
			token.End(),
		}),
	},
	{
		input: "(add 1 2)",
		output: token.NewList([]token.Token{
			token.Start(),
			token.Statement("add"),
			token.Int("1"),
			token.Int("2"),
			token.End(),
		}),
	},
	{
		input: "(if true false )", // TODO: why is this space needed?
		output: token.NewList([]token.Token{
			token.Start(),
			token.Statement("if"),
			token.True(),
			token.False(),
			token.End(),
		}),
	},
}

func TestLexerTokens(t *testing.T) {
	for _, tst := range lexerTests {
		t.Run("", func(t *testing.T) {
			l := Lexer{}
			l.WriteString(tst.input)
			expectEqual(t, l.tokens.Length(), tst.output.Length())

			actual := l.Tokens()
			expected := tst.output
			if slices.EqualFunc(actual.Tokens, expected.Tokens, token.Compare) {
				t.Logf("%v == %v", actual, expected)
			} else {
				t.Errorf("%v != %v", actual, expected)
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
	actual := l.Tokens().Tokens[0]
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
	requireEqual(t, l.tokens.Length(), 1)
	actual := l.Tokens().Tokens[0]
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
