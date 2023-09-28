package ast

import (
	"fmt"
	"slices"
)

// All AST elements implement the Element interface
type Element interface {
	Type() Type
}

type Type int

const (
	VOID = iota
	LIST
	INT
	BOOL
	STRING
	IDENTIFIER
	CALL
)

var typeString = map[int]string{
	VOID:       "VOID",
	LIST:       "LIST",
	INT:        "INT",
	BOOL:       "BOOL",
	STRING:     "STRING",
	IDENTIFIER: "IDENTIFIER",
	CALL:       "CALL",
}

type Void int

func (v Void) Type() Type {
	return VOID
}
func NewVoid() Void {
	return Void(0)
}

type Int int

func (i Int) Type() Type {
	return INT
}

type Bool bool

func (Bool) Type() Type {
	return BOOL
}

func True() Bool {
	return true
}

func False() Bool {
	return false
}

type String string

func (s String) Type() Type {
	return STRING
}

func (s String) String() string {
	return fmt.Sprintf("\"%s\"", string(s))
}

type Identifier string

func (i Identifier) Type() Type {
	return IDENTIFIER
}

type Call struct {
	Name       Identifier
	Parameters List
}

func (c Call) Type() Type {
	return CALL
}

func (c Call) String() string {
	return fmt.Sprintf("{fn %s, parameters %v}", c.Name, c.Parameters)
}

func Equal(a, b Element) bool {
	// compare types
	if a.Type() != b.Type() {
		return false
	}

	// compare lists
	if a, ok := a.(List); ok {
		b := b.(List)
		return slices.EqualFunc(a, b, Equal)
	}

	// compare function calls
	if a, ok := a.(Call); ok {
		b := b.(Call)
		if a.Name != b.Name {
			return false
		}
		return Equal(a.Parameters, b.Parameters)
	}

	// compare generic
	return a == b
}
