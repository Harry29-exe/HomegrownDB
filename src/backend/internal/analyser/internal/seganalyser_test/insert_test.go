package seganalyser_test

import (
	"HomegrownDB/backend/internal/analyser/internal/seganalyser"
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/common/bparse"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tstructs"
	ttable12 "HomegrownDB/common/tests/tutils/testtable/ttable1"
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
		table1 := ttable12.Def(t)
		tableStore := tstructs.NewTestTableStoreWithInMemoryIO(t, table1)
		txCtx := tx.NewContext(25, tableStore)

		query := fmt.Sprintf(query, ttable12.TableName, ttable12.C0AwesomeKey, ttable12.C2NonNullColl, ttable12.C1NullableCol)
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
		assert.Eq(insertANode.Columns[0].Id(), table1.Column(ttable12.C0AwesomeKeyOrder).Id(), t)
		assert.Eq(insertANode.Columns[1].Id(), table1.Column(ttable12.C2NonNullCollOrder).Id(), t)
		assert.Eq(insertANode.Columns[2].Id(), table1.Column(ttable12.C1NullableColOrder).Id(), t)

		rows := insertANode.Rows
		assert.EqArray(rows.GetValue(0, 0), bparse.Serialize.Int8(2), t)
		assert.EqArray(rows.GetValue(0, 1), bparse.Serialize.Int8(3), t)
		assert.EqArray(rows.GetValue(0, 2), bparse.Serialize.Int8(4), t)
		assert.EqArray(rows.GetValue(1, 0), bparse.Serialize.Int8(5), t)
		assert.EqArray(rows.GetValue(1, 1), bparse.Serialize.Int8(6), t)
		assert.EqArray(rows.GetValue(1, 2), bparse.Serialize.Int8(7), t)
	}
}
