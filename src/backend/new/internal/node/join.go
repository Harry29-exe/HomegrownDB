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

func (f fromExpr) DEqual() bool {
	//TODO implement me
	panic("implement me")
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
