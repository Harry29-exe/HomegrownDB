package exenode

import (
	"HomegrownDB/dbsystem/access"
	dbbs2 "HomegrownDB/dbsystem/access/dbbs"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
)

type Insert struct {
	table   table.Definition
	tableIO access.TableDataIO
	rowSrc  InputRowSrc

	txCtx          tx.Ctx
	valuesInserted bool
}

func (i *Insert) SetSource(source []ExeNode) {
	panic("operation not supported: leaf exe node")
}

func (i *Insert) HasNext() bool {
	return !i.valuesInserted
}

// todo(3) rebuild this properly
func (i *Insert) Next() dbbs2.QRow {
	tupleData := i.rowSrc.NextRow()
	tuple := page.NewTuple(tupleData, i.table, i.txCtx)

	return dbbs2.NewQRowFromTuple(tuple)
}

func (i *Insert) NextBatch() []dbbs2.QRow {
	//TODO implement me
	panic("implement me")
}

func (i *Insert) All() []dbbs2.QRow {
	//TODO implement me
	panic("implement me")
}

func (i *Insert) Free() {
	//TODO implement me
	panic("implement me")
}
