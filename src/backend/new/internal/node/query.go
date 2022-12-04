package node

func NewQuery() Query {
	//todo implement me
	panic("Not implemented")
}

type Query = *query

type query struct {
	node
	Command    CommandType
	TargetList []TargetEntry
	RTables    []RangeTableEntry
	FromExpr   FromExpr
}

// -------------------------
//      FromExpr
// -------------------------

func NewFromExpr() FromExpr {
	return &fromExpr{
		node:       node{tag: TagFrom},
		FromList:   make([]Node, 0, 8),
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
