package node

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

func (f FromExpr) DEqual(node Node) bool {
	if res, ok := nodeEqual(f, node); ok {
		return res
	}
	raw := node.(FromExpr)
	return cmpNodeArray(f.FromList, raw.FromList) &&
		f.Qualifiers.DEqual(raw.Qualifiers)
}

// -------------------------
//      JoinType
// -------------------------

type JoinType = uint8

const (
	JoinInner JoinType = iota
	JoinLeft
	JoinFull
	JoinRight
)
