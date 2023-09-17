package lexer

import (
	"fmt"
	"testing"

	"github.com/monban/lispian/token"
)

var lexerTests = [...]lexerTest{
	{
		text:       "()",
		tokenCount: 2,
		tokens:     []token.Token{token.Start(), token.End()},
	},
	{
		text:       "(\"Hello, world\")",
		tokenCount: 3,
		tokens:     []token.Token{token.Start(), token.String("Hello, world"), token.End()},
	},
}

func TestLexerCount(t *testing.T) {
	for _, tst := range lexerTests {
		t.Run("", func(t *testing.T) {
			l := Lexer{}
			l.Write([]byte(tst.text))
			expectEqual(t, tst.tokenCount, len(l.tokens))
		})
	}
}
func TestLexerTokens(t *testing.T) {
	for _, tst := range lexerTests {
		t.Run("", func(t *testing.T) {
			l := Lexer{}
			l.Write([]byte(tst.text))
			if token.CompareSlice(tst.tokens, l.tokens) == true {
				t.Logf("%s == %s", tst.tokens, l.tokens)
			} else {
				t.Errorf("%s != %s", tst.tokens, l.tokens)
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

type lexerTest struct {
	text       string
	tokenCount int
	tokens     []token.Token
}

func (lt lexerTest) String() string {
	return fmt.Sprintf("%s %d", lt.text, lt.tokenCount)
}
