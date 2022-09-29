package seganalyser_test

import (
	"HomegrownDB/backend/internal/analyser/internal/seganalyser"
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tstructs"
	"HomegrownDB/common/tests/tutils"
	"HomegrownDB/dbsystem/tx"
	"fmt"
	"testing"
)

func TestInsertSimpleQuery(t *testing.T) {
	queries := []string{
		"INSERT INTO %s (%s, %s, %s) VALUES (2, 3, 4), (5, 6, 7)",
		"INSERT INTO %s ( %s , %s , %s ) VALUES ( 2, 3, 4 ) , ( 5, 6, 7 )",
		"INSERT INTO %s (%s,%s,%s) VALUES (2,3,4),(5,6,7)",
		"INSERT INTO %s ( %s ,  %s  ,%s) VALUES (  2, 3, 4),\n ( 5,   6 ,7)",
	}

	for _, query := range queries {
		// given
		table1 := tutils.TestTables.Table1Def()
		tableStore := tstructs.NewTestTableStoreWithInMemoryIO(table1)
		txCtx := tx.NewContext(25, tableStore)

		query := fmt.Sprintf(query, tutils.Table1Name, tutils.Table1AwesomeKey, tutils.Table1NonNullColl, tutils.Table1NullableCol)
		parserTree, err := parser.Parse(query, txCtx)
		if err != nil {
			t.Error(err)
		}
		insertPnode := parserTree.Root.(pnode.InsertNode)

		// when
		insertANode, err := seganalyser.Insert.Analyse(insertPnode, txCtx)

		//then
		if err != nil {
			t.Error(err)
		}
		table := insertANode.Table
		tableId := table.Def.TableId()
		assert.Eq(table.Def.Name(), table1.Name(), t)
		assert.Eq(table.Alias, table1.Name(), t)
		assert.Eq(tableId, table1.TableId(), t)

		assert.Eq(len(insertANode.Columns), 3, t)
		assert.Eq(insertANode.Columns[0].Id(), table1.Column(tutils.Table1AwesomeKeyId).Id(), t)
		assert.Eq(insertANode.Columns[1].Id(), table1.Column(tutils.Table1NonNullCollId).Id(), t)
		assert.Eq(insertANode.Columns[2].Id(), table1.Column(tutils.Table1NullableColId).Id(), t)

		rows := insertANode.Rows
		assert.EqArray(rows.GetValue(0, 0), bparse.Serialize.Int8(2), t)
		assert.EqArray(rows.GetValue(0, 1), bparse.Serialize.Int8(3), t)
		assert.EqArray(rows.GetValue(0, 2), bparse.Serialize.Int8(4), t)
		assert.EqArray(rows.GetValue(1, 0), bparse.Serialize.Int8(5), t)
		assert.EqArray(rows.GetValue(1, 1), bparse.Serialize.Int8(6), t)
		assert.EqArray(rows.GetValue(1, 2), bparse.Serialize.Int8(7), t)
	}
}
