package ast

import (
	"fmt"
	"strings"
)

type List []Element

func (l List) Type() Type {
	return LIST
}

func (l *List) AddElement(e Element) {
	*l = append(*l, e)
}

func (l List) String() string {
	return l.StringIndent(0)
}

func (l List) StringIndent(depth int) string {
	b := &strings.Builder{}
	for i := 0; i < depth; i++ {
		b.WriteString("\t")
	}
	fmt.Fprintf(b, "(\n")
	for _, e := range l {
		if l, ok := e.(List); ok {
			b.WriteString(l.StringIndent(depth + 1))
		} else {
			for i := 0; i < depth; i++ {
				b.WriteString("\t")
			}
			b.WriteString("\t")
			fmt.Fprintf(b, "%v\n", e)
		}
	}
	for i := 0; i < depth; i++ {
		b.WriteString("\t")
	}
	fmt.Fprintf(b, ")\n")
	return b.String()
}
