package main

import (
	"fmt"
	"io"
	"os"

	"github.com/monban/lispian/lexer"
)

func main() {
	f, err := os.Open("text")
	p := &lexer.Lexer{}
	if err != nil {
		panic(err)
	}
	io.Copy(p, f)
	fmt.Println(p.Tokens())
}
