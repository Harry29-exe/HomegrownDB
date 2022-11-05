package exenode

import (
	"HomegrownDB/dbsystem/access/dbbs"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/tx"
)

var _ ExeNode = &Insert{}

type Insert struct {
	table table.Definition

	rowSrc ExeNode

	txCtx          *tx.Ctx
	valuesInserted bool
}

func (i *Insert) SetSource(source []ExeNode) {
	panic("operation not supported: leaf exe node")
}

func (i *Insert) HasNext() bool {
	return !i.valuesInserted
}

// todo(3) rebuild this properly
func (i *Insert) Next() dbbs.QRow {
	//todo implement me
	panic("Not implemented")
	//tupleData := i.rowSrc.Next()
	//tuple := tpage.NewTuple(tupleData, i.table, i.txCtx)
	//
	//return dbbs.NewQRowFromTuple(tuple)
}

func (i *Insert) NextBatch() []dbbs.QRow {
	//TODO implement me
	panic("implement me")
}

func (i *Insert) All() []dbbs.QRow {
	//TODO implement me
	panic("implement me")
}

func (i *Insert) Free() {
	//TODO implement me
	panic("implement me")
}
