package plan

import (
	"HomegrownDB/dbsystem/schema/table"
)

type SeqScan struct {
	Table      table.Id
	Conditions Conditions
}

func (s SeqScan) Type() NodeType {
	return SeqScanNode
}

func (s SeqScan) Children() []Node {
	return nil
}
