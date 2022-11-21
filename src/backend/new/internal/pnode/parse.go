package node

type SelectStmt struct {
	PNode
	Targets []ResultTarget
	From    []RangeVar
	Where   AExpr
}
