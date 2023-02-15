package pnode

func NewCommandStmt(stmt Node) CommandStmt {
	return &commandStmt{
		node: node{
			tag: TagCommandStmt,
		},
		Stmt: stmt,
	}
}

type CommandStmt = *commandStmt

var _ Node = &commandStmt{}

type commandStmt struct {
	node
	Stmt Node
}

func (c CommandStmt) Equal(node Node) bool {
	if nodesEqNil(c, node) {
		return true
	} else if !basicNodeEqual(c, node) {
		return false
	}

	raw := node.(CommandStmt)
	return c.Stmt.Equal(raw)
}

// -------------------------
//      CreateTableStmt
// -------------------------

func NewCreateTableStmt(tableName string, columns []ColumnDef) CreateTableStmt {
	return &createTableStmt{
		node: node{
			tag: TagCreateTable,
		},
		TableName: tableName,
		Columns:   columns,
	}
}

type CreateTableStmt = *createTableStmt

var _ Node = &createTableStmt{}

type createTableStmt struct {
	node
	TableName string
	Columns   []ColumnDef
}

func (c CreateTableStmt) Equal(node Node) bool {
	if nodesEqNil(c, node) {
		return true
	} else if !basicNodeEqual(c, node) {
		return false
	}
	raw := node.(CreateTableStmt)
	return c.TableName == raw.TableName &&
		cmpNodeArray(c.Columns, raw.Columns)
}
