package parser

type ParserError error

type UnbalancedParams struct {
	error
}
