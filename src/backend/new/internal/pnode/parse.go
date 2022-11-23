package pnode

type SelectStmt = *selectStmt

func NewSelectStmt() SelectStmt {
	return &selectStmt{
		node: node{tag: TagSelectStmt},
	}
}

type selectStmt struct {
	node
	Targets []ResultTarget
	From    []RangeVar
	Where   AExpr
}
