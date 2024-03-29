package table

import (
	"HomegrownDB/dbsystem/access/buffer"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/storage/fsm"
	"HomegrownDB/dbsystem/storage/page"
	"HomegrownDB/dbsystem/tx"
)

func Insert(
	tuple page.WTuple,
	tx tx.Tx,
	table reldef.TableRDefinition,
	fsm *fsm.FSM,
	buffer buffer.SharedBuffer,
) error {
	pageId, err := fsm.FindPage(uint16(tuple.TupleSize()), tx)
	if err != nil {
		panic(err.Error())
	}
	wPage, err := buffer.WTablePage(table, pageId)
	if err != nil {
		panic(err.Error())
	}
	defer buffer.WPageRelease(wPage.PageTag())

	return wPage.InsertTuple(tuple.Bytes())
}
