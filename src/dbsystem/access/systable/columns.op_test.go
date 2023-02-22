package systable

import (
	"HomegrownDB/dbsystem/access/systable/internal"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/lib/tests/assert"
	"testing"
)

func Test_ColumnsOps_SerializeDeserialize(t *testing.T) {
	// given
	tableOID := reldef.OID(119)
	col := reldef.NewColumnDefinition("col_name1", reldef.OID(124), reldef.Order(2), hgtype.NewInt8(hgtype.Args{Nullable: true}))
	testTx := tx.StdTx{
		Id:                52,
		TxLevel:           0,
		TxStatus:          0,
		TxCommandExecuted: 2,
	}

	// when
	tuple := ColumnsOps.DataToRow(tableOID, col, testTx)
	newCol := ColumnsOps.RowToData(tuple)

	//then
	assert.Eq(col.Name(), newCol.Name(), t)
	assert.Eq(col.Id(), newCol.Id(), t)
	assert.Eq(col.Order(), newCol.Order(), t)
	colType, newColType := col.CType(), newCol.CType()
	assert.Eq(colType.Tag(), newColType.Tag(), t)
	assert.Eq(colType.Args(), newColType.Args(), t)

	tableOIDFromTuple := internal.GetInt8(ColumnsOrderRelationOID, tuple)
	assert.Eq(tableOID, reldef.OID(tableOIDFromTuple), t)
}
