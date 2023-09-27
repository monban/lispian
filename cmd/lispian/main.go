package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/monban/lispian/interpreter"
	"github.com/monban/lispian/lexer"
	"github.com/monban/lispian/parser"
)

func main() {
	fmt.Println("Welcome to the Lispian REPL!")
	scanner := bufio.NewScanner(os.Stdin)
	l := lexer.Lexer{}
	for {
		fmt.Printf("> ")
		scanner.Scan()
		input := scanner.Text()
		if input == "quit" {
			break
		}
		if input[0] != '(' {
			input = "(" + input
		}
		if input[len(input)-1] != ')' {
			input = input + ")"
		}
		fmt.Printf("DEBUG: input: %s\n", input)
		err := l.WriteString(input)
		if err != nil {
			panic(err)
		}
		fmt.Printf("DEBUG: tokens: %#v\n", l.Tokens())

		program, err := parser.Parse(l.Tokens())
		if err != nil {
			panic(err)
		}
		fmt.Printf("DEBUG: program: %#v\n", l)

		output := interpreter.Eval(program)
		fmt.Printf("%v\n", output)
		l.Reset()
	}
}
