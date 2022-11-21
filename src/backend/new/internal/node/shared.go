package node

import "HomegrownDB/dbsystem/schema/table"

type CommandType = uint8

const (
	// CommandTypeSelect for select statements
	CommandTypeSelect CommandType = iota
	// CommandTypeInsert for insert statements
	CommandTypeInsert
	// CommandTypeUpdate for update statements
	CommandTypeUpdate
	// CommandTypeDelete for delete statements
	CommandTypeDelete
	// CommandTypeUtils for statements that operate
	// rater on db structure than on db data
	CommandTypeUtils
)

// RangeTableEntry is db table that is used in query or plan
type RangeTableEntry struct {
	Node
	Id       RteID
	LockMode table.TableLockMode
	TableId  table.Id
	Ref      table.Definition
}

// RteID is id of RangeTableEntry unique for given query/plan
type RteID = uint16

// RangeTableRef is ref to RTE in query/plan
type RangeTableRef struct {
	Node
	Rte RteID
}

type TargetEntry struct {
	Expr                // Expr to treat TargetEntry as Expr node
	ExprToExec *Expr    // ExprToExec expression to evaluate to
	AttribNo   uint16   // AttribNo number of entry
	TableId    table.Id // TableId
	ColName    string   // ColName nullable column name
	Temp       bool     // Temp if true then entry should be eliminated before tuple is emitted
}
