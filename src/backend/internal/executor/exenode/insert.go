package exenode

import (
	"HomegrownDB/backend/internal/shared/query"
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/table"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/tpage"
	"HomegrownDB/dbsystem/tx"
)

var _ ExeNode = &Insert{}

type Insert struct {
	table   table.Definition
	columns []column.Id

	rowSrc ExeNode

	buff  buffer.SharedBuffer
	fsm   *fsm.FreeSpaceMap
	txCtx *tx.Ctx

	// tupleTemplate is array of values, initialized before
	// node execution to not allocation every time new array
	tupleTemplate [][]byte
	// colOrder is map[index of data in QRow from source] = column order in table
	colOrder []int
}

func (i *Insert) SetSource(source []ExeNode) {
	panic("operation not supported: leaf exe node")
}

func (i *Insert) HasNext() bool {
	return i.rowSrc.HasNext()
}

func (i *Insert) Next() query.QRow {
	tupleData := i.rowSrc.Next()
	for j := 0; j < len(i.colOrder); j++ {
		order := i.colOrder[j]
		i.tupleTemplate[order] = tupleData.Value(uint32(j))
	}

	tuple := tpage.NewTuple(i.tupleTemplate, i.table, i.txCtx)
	insertSuccessful := false
	for !insertSuccessful {
		wPage := i.findPage(uint16(len(tuple.Data())))
		err := wPage.InsertTuple(tuple.Data())
		if err == nil {
			insertSuccessful = true
		}
		i.buff.WPageRelease(wPage.PageTag())
	}

	return tupleData
}

func (i *Insert) NextBatch() []query.QRow {
	//todo this is definetly not optimized
	return []query.QRow{i.Next()}
}

func (i *Insert) All() []query.QRow {
	rows := make([]query.QRow, 50)
	for i.rowSrc.HasNext() {
		rows = append(rows, i.NextBatch()...)
	}
	return rows
}

func (i *Insert) Free() {
	//TODO implement me
	panic("implement me")
}

func (i *Insert) insertIntoNewPage(tuple tpage.Tuple) {

}

func (i *Insert) findPage(neededSpace uint16) tpage.WPage {
	pageId, err := i.fsm.FindPage(neededSpace, i.txCtx)
	if err != nil {
		panic(err.Error())
	}

	tablePage, err := i.buff.WTablePage(i.table, pageId)
	if err != nil {
		panic(err.Error())
	}

	return tablePage
}
