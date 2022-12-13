package plan

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
)

var _ Node = &insert{}

type Insert = *insert

func NewInsert(table table.Id, columns []column.Order) Insert {
	return &insert{
		Table:   table,
		Columns: columns,
		Src:     nil,
	}
}

type insert struct {
	Table   table.Id
	Columns []column.Order
	Src     Node
}

func (i Insert) Type() NodeType {
	//TODO implement me
	panic("implement me")
}

func (i Insert) Children() []Node {
	//TODO implement me
	panic("implement me")
}
