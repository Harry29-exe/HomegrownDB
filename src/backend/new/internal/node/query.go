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

type FromExpr struct {
}
