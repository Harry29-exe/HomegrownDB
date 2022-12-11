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
	Stmt Node
	//todo implement StmtSource so it can be used to show where error is in
	// query, it should be based on tokenizer.TokenSource
	TkSource StmtSource // TkSource for error/debugging purposes
}

func (r *rawStmt) Equal(node Node) bool {
	if nodesEqNil(r, node) {
		return true
	} else if !basicNodeEqual(r, node) {
		return false
	}
	raw := node.(RawStmt)
	return r.Stmt.Equal(raw.Stmt)
}

type StmtSource interface {
	History() []any
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

func NewSelectStmtWithValues(values [][]Node) SelectStmt {
	return &selectStmt{
		node:   node{tag: TagSelectStmt},
		Values: values,
	}
}

var _ Node = &selectStmt{}

type selectStmt struct {
	node
	Targets []ResultTarget
	From    []Node // list of targets, it can be list of
	Where   AExpr

	Values [][]Node // values for value select (A_Const/A_Expr/FuncCall/
}

func (s SelectStmt) Equal(node Node) bool {
	if nodesEqNil(s, node) {
		return true
	} else if !basicNodeEqual(s, node) {
		return false
	}
	raw := node.(SelectStmt)
	return cmpNodeArray(s.Targets, raw.Targets) &&
		cmpNodeArray(s.From, raw.From) &&
		s.Where.Equal(raw.Where) &&
		cmpNodeArray2D(s.Values, raw.Values)
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
	if nodesEqNil(s, node) {
		return true
	} else if !basicNodeEqual(s, node) {
		return false
	}
	raw := node.(InsertStmt)

	return cmpNodeArray(s.Columns, raw.Columns) &&
		s.Relation.Equal(raw.Relation) &&
		s.SrcNode.Equal(raw.SrcNode)
}
