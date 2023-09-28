package parser

import (
	"fmt"
	"strconv"

	"github.com/monban/lispian/ast"
	"github.com/monban/lispian/token"
)

type Parser struct {
	program ast.List
}

func (p *Parser) Parse(ts []token.Token) (ast.Element, error) {
	for i := 0; i < len(ts); i++ {
		value := ts[i].Text
		switch ts[i].Type {
		case token.STATEMENT:
			p.program.AddElement(ast.Identifier(value))
		case token.STRING:
			p.program.AddElement(ast.String(value))
		case token.INT:
			integer, _ := strconv.ParseInt(value, 10, 32)
			p.program.AddElement(ast.Int(integer))
		case token.LIST_START:
			sublist, j, err := p.parseList(ts[i:])
			if err != nil {
				return ast.List{}, fmt.Errorf("error parsing %v: %w", ts[i], err)
			}
			i += j
			p.program.AddElement(sublist)
		case token.BOOL:
			if value == "true" {
				p.program.AddElement(ast.True())
			} else if value == "false" {
				p.program.AddElement(ast.False())
			} else {
				panic("invalid boolean")
			}
		default:
			return ast.List{}, fmt.Errorf("error parsing %v", ts[i])
		}
	}
	return p.program[0], nil
}

func (p *Parser) parseList(ts []token.Token) (ast.Element, int, error) {
	fmt.Println("parsing sublist: ", ts)
	i := 1
	var out ast.Element
	if ts[i].Type == token.STATEMENT {
		// This is a function call
		call := ast.Call{
			Name: ast.Identifier(ts[i].Text),
		}
		for ; i < len(ts); i++ {
			if ts[i].Type == token.LIST_END {
				return call, i, nil
			}
			e, _ := p.Parse(ts[i:])
			call.Parameters.AddElement(e)
		}
		fmt.Println(call)
		out = call
	} else {
		l := ast.List{}
		for ; i < len(ts); i++ {
			if ts[i].Type == token.LIST_END {
				break
			}
			element, _ := p.Parse(ts[i:])
			l.AddElement(element)
		}
		out = l
	}
	fmt.Println("returning sublist: ", out)
	return out, i, nil
}
