package node

func NewQuery() Query {
	//todo implement me
	panic("Not implemented")
}

type Query = *query

type query struct {
	Node
	Command    CommandType
	targetList []TargetEntry
	rTables    []RangeTableEntry
	fromExpr   FromExpr
}

type FromExpr struct {
}
