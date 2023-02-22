package systable

import (
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
)

var SequencesOps = sequencesOps{}

type sequencesOps struct{}

func (sequencesOps) DataToRow(sequence reldef.SequenceDef, tx tx.Tx) page.WTuple {
	//builder := internal.NewTupleBuilder(sequenceDef)
	//todo implement me
	panic("Not implemented")
}
