package pnode

type CommandStmt interface {
	Node
	CommandType() CommandType
}

type commandStmt struct {
	node
	Type CommandType
}

func (c *commandStmt) CommandType() CommandType {
	return c.Type
}

type CommandType uint8

const (
	CommandCreateTable CommandType = iota
)

// -------------------------
//      CreateTableStmt
// -------------------------

func NewCreateTableStmt(tableName string, columns []ColumnDef) CreateTableStmt {
	return &createTableStmt{
		commandStmt: commandStmt{
			node: node{
				tag:        0,
			},
			Type: CommandCreateTable,
		},
		TableName:   tableName,
		Columns:     columns,
	}
}

type CreateTableStmt = *createTableStmt

var _ CommandStmt = &createTableStmt{}

type createTableStmt struct {
	commandStmt
	TableName string
	Columns []ColumnDef
}

func (c createTableStmt) Equal(node Node) bool {
	//TODO implement me
	panic("implement me")
}
