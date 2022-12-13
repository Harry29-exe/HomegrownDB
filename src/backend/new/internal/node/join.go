package node

import "fmt"

// -------------------------
//      FromExpr
// -------------------------

func NewFromExpr(fromListSize int) FromExpr {
	return &fromExpr{
		node:       node{tag: TagFrom},
		FromList:   make([]Node, fromListSize),
		Qualifiers: nil,
	}
}

type FromExpr = *fromExpr

var _ Node = &fromExpr{}

type fromExpr struct {
	node
	FromList   []Node
	Qualifiers Node
}

func (f FromExpr) dEqual(node Node) bool {
	raw := node.(FromExpr)
	return cmpNodeArray(f.FromList, raw.FromList) &&
		DEqual(f.Qualifiers, raw.Qualifiers)
}

func (f FromExpr) DPrint(nesting int) string {
	inner := fmt.Sprintf("{\nFromList: %s\nQualifiers: %s\n}",
		dPrintArr(nesting+1, f.FromList),
		f.Qualifiers.DPrint(nesting+1),
	)
	return f.dFormat(inner, nesting)
}

// -------------------------
//      JoinType
// -------------------------

type JoinType uint8

const (
	JoinInner JoinType = iota
	JoinLeft
	JoinFull
	JoinRight
)

func (j JoinType) ToString() string {
	return [...]string{
		"JoinInner",
		"JoinLeft",
		"JoinFull",
		"JoinRight",
	}[j]
}
