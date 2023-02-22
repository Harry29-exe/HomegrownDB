package systable

import (
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/lib/tests/assert"
	"reflect"
	"testing"
)

func TestRelationsOps_SerializeDeserialize(t *testing.T) {
	// given
	oldTable := reldef.NewTableDefinition("super_table")
	oldTable.InitRel(56, 58, 66)
	err := oldTable.AddColumn(reldef.NewColumnDefinition(
		"col1", 67, 0, hgtype.NewInt8(hgtype.Args{Nullable: true})))
	assert.ErrIsNil(err, t)
	err = oldTable.AddColumn(reldef.NewColumnDefinition(
		"col2", 69, 1, hgtype.NewStr(hgtype.Args{UTF8: true, VarLen: true})))
	assert.ErrIsNil(err, t)

	currentTx := tx.StdTx{
		Id:                623,
		TxCommandExecuted: 1,
	}

	// when
	tableTuple, err := RelationsOps.TableAsRelationsRow(oldTable, currentTx)
	assert.ErrIsNil(err, t)
	columnTuples := ColumnsOps.DataToRows(oldTable.OID(), oldTable.Columns(), currentTx)

	newTable := RelationsOps.RowAsData(tableTuple).(reldef.TableDefinition)
	for _, colTuple := range columnTuples {
		colDef := ColumnsOps.RowToData(colTuple)
		err = newTable.AddColumn(colDef)
		assert.ErrIsNil(err, t)
	}

	// then
	assert.True(reflect.DeepEqual(oldTable, newTable), t)
}
