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

// todo(3) rebuild this properly
func (i *Insert) Next() query.QRow {
	tupleData := i.rowSrc.Next()
	for j := 0; j < len(i.colOrder); j++ {
		order := i.colOrder[j]
		i.tupleTemplate[order] = tupleData.Value(uint32(j))
	}

	//todo implement me
	panic("Not implemented")
	//tuple := tpage.NewTuple(i.tupleTemplate, i.table, i.txCtx)
	//freePageIndex, err := i.fsm.FindPage(uint16(len(tuple.Data())), i.txCtx)
	//if err != nil {

	//}

	//return query.NewQRowFromValues([][]byte{bparse.Serialize.Int8(1)}, []ctype.CType{ctype.TypeInt8})
}

func (i *Insert) NextBatch() []query.QRow {
	//TODO implement me
	panic("implement me")
}

func (i *Insert) All() []query.QRow {
	//TODO implement me
	panic("implement me")
}

func (i *Insert) Free() {
	//TODO implement me
	panic("implement me")
}

func (i *Insert) insertIntoNewPage(tuple tpage.Tuple) {

}
