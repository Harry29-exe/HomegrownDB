package pnode

// -------------------------
//      RawStmt
// -------------------------

type RawStmt = *rawStmt

func NewRawStmt() RawStmt {
	return &rawStmt{
		node: node{tag: TagRawStmt},
	}
}

var _ Node = &rawStmt{}

type rawStmt struct {
	node
	stmt Node
}

func (r *rawStmt) Equal(node Node) bool {
	if !basicNodeEqual(r, node) {
		return false
	}
	raw := node.(RawStmt)
	return r.stmt.Equal(raw.stmt)
}

// -------------------------
//      SelectStmt
// -------------------------

// SelectStmt represents stream of values, usually eiter statements
// starting from SELECT token or VALUES token
type SelectStmt = *selectStmt

func NewSelectStmt() SelectStmt {
	return &selectStmt{
		node: node{tag: TagSelectStmt},
	}
}

var _ Node = &selectStmt{}

type selectStmt struct {
	node
	Targets []ResultTarget
	From    []RangeVar
	Where   AExpr

	Values [][]Node // values for value select (A_Const/A_Expr/FuncCall/
}

func (s SelectStmt) Equal(node Node) bool {
	//todo implement me
	panic("Not implemented")
}

// -------------------------
//      InsertStmt
// -------------------------

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

func (s InsertStmt) Equal(node Node) bool {
	//todo implement me
	panic("Not implemented")
}
