package exenode

import (
	"HomegrownDB/dbsystem/access"
	"HomegrownDB/dbsystem/dbbs"
	"HomegrownDB/dbsystem/schema/table"
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
func (i *Insert) Next() dbbs.QRow {
	//tuple := dbbs.NewTestTuple()
	//todo implement me
	panic("Not implemented")
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
