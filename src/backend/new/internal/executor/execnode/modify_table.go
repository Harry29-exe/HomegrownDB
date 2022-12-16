package execnode

import (
	"HomegrownDB/backend/new/internal/node"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/storage/tpage"
)

var _ ExecNode = ModifyTable{}

type ModifyTable struct {
	plan node.Plan
	buff buffer.SharedBuffer
}

func (m ModifyTable) Next() []tpage.Tuple {
	//TODO implement me
	panic("implement me")
}

func (m ModifyTable) Init(plan node.Plan) error {
	//TODO implement me
	panic("implement me")
}
