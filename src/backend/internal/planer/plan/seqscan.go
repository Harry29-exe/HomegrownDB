package plan

import (
	"HomegrownDB/dbsystem/schema/table"
)

func NewSeqScan(def table.RDefinition) SeqScan {
	return &seqScan{
		Table:      Table{TableId: def.TableId(), TableHash: def.Hash()},
		Conditions: nil,
	}
}

type SeqScan = *seqScan

type seqScan struct {
	Table      Table
	Conditions Conditions
}

func (s seqScan) Type() NodeType {
	return SeqScanNode
}

func (s seqScan) Children() []Node {
	return nil
}
