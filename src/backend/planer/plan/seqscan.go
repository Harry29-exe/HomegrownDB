package plan

import (
	"HomegrownDB/dbsystem/schema/table"
)

type SeqScan struct {
	Table      table.Definition
	Conditions Conditions
}

func (s SeqScan) Type() NodeType {
	return SeqScanNode
}

func (s SeqScan) Children() []Node {
	return nil
}
