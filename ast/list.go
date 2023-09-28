package ast

type List []Element

func (l List) Type() Type {
	return LIST
}

func (l *List) AddElement(e Element) {
	*l = append(*l, e)
}
