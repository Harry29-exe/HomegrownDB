package seganalyser_test

import (
	"HomegrownDB/backend/internal/analyser/internal/seganalyser"
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/backend/internal/shared/qctx"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/ttable1"
	"HomegrownDB/dbsystem/schema/table"
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
		table1 := ttable1.Def(t)
		tableStore, err := table.NewTableStore([]table.Definition{table1})
		assert.IsNil(err, t)
		qCtx := qctx.NewQueryCtx(tableStore)

		query := fmt.Sprintf(query, ttable1.TableName, ttable1.C0AwesomeKey, ttable1.C2NonNullColl, ttable1.C1NullableCol)
		parserTree, err := parser.Parse(query, qCtx)
		if err != nil {
			t.Error(err)
		}
		insertPnode := parserTree.Root.(pnode.InsertNode)

		// when
		insertANode, err := seganalyser.Insert.Analyse(insertPnode, qCtx)

		//then
		if err != nil {
			t.Error(err)
		}
		tableDef := insertANode.Table
		println(tableDef)
		//tableId := tableDef.Def.TableId()
		//assert.Eq(tableDef.Def.Name(), table1.Name(), t)
		//assert.Eq(tableDef.Alias, table1.Name(), t)
		//assert.Eq(tableId, table1.TableId(), t)
		//
		//assert.Eq(len(insertANode.Columns), 3, t)
		//assert.Eq(insertANode.Columns[0].Id(), table1.Column(ttable1.C0AwesomeKeyOrder).Id(), t)
		//assert.Eq(insertANode.Columns[1].Id(), table1.Column(ttable1.C2NonNullCollOrder).Id(), t)
		//assert.Eq(insertANode.Columns[2].Id(), table1.Column(ttable1.C1NullableColOrder).Id(), t)
		//
		//rows := insertANode.Rows
		//assert.EqArray(rows[0].Fields[0].Value, bparse.Serialize.Int8(2), t)
		//assert.EqArray(rows[0].Fields[1].Value, bparse.Serialize.Int8(3), t)
		//assert.EqArray(rows[0].Fields[2].Value, bparse.Serialize.Int8(4), t)
		//assert.EqArray(rows[1].Fields[0].Value, bparse.Serialize.Int8(5), t)
		//assert.EqArray(rows[1].Fields[1].Value, bparse.Serialize.Int8(6), t)
		//assert.EqArray(rows[1].Fields[2].Value, bparse.Serialize.Int8(7), t)
	}
}
