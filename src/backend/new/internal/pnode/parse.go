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

	Values [][]Node // values for value select (A_Const/A_Expr/FuncCall/
}

type InsertStmt = *insertStmt

func NewInsertStmt() InsertStmt {
	return &insertStmt{
		node: node{tag: TagInsertStmt},
	}
}

type insertStmt struct {
	node
	Relation RangeVar // Relation that rows will be inserted to
	Columns  []ResultTarget
	SrcNode  Node // source of nodes to be inserted
}
