package plan

import (
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
)

var _ Node = &Insert{}

type Insert struct {
	Table   table.Definition
	Columns []column.Def
	Src     Node
}

func (i *Insert) Type() NodeType {
	//TODO implement me
	panic("implement me")
}

func (i *Insert) Children() []Node {
	//TODO implement me
	panic("implement me")
}
