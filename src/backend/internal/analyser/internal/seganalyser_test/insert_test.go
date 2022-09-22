package seganalyser_test

import (
	"HomegrownDB/backend/internal/analyser/internal"
	"HomegrownDB/backend/internal/analyser/internal/seganalyser"
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tstructs"
	"HomegrownDB/common/tests/tutils"
	"HomegrownDB/dbsystem/tx"
	"fmt"
	"testing"
)

func TestInsertSimpleQuery(t *testing.T) {
	// given
	table1 := tutils.TestTables.Table1Def()
	tableStore := tstructs.NewTestTableStoreWithInMemoryIO(table1)
	txCtx := tx.NewContext(25)

	query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s) VALUES (2, 3, 4), (5, 6, 7)",
		tutils.Table1Name, tutils.Table1AwesomeKey, tutils.Table1NonNullColl, tutils.Table1NullableCol)
	parserTree, err := parser.Parse(query, txCtx)
	if err != nil {
		t.Error(err)
	}
	insertPnode := parserTree.Root.(pnode.InsertNode)

	// when
	aCtx := internal.NewAnalyserCtx(txCtx, tableStore)
	insertANode, err := seganalyser.Insert.Analyse(insertPnode, aCtx)
	if err != nil {
		t.Error(err)
	}

	table := insertANode.Table
	tableId := table.Def.TableId()
	assert.Eq(table.Def.Name(), "users", t)
	assert.Eq(table.Alias, "users", t)
	assert.Eq(tableId, table1.TableId(), t)

	assert.Eq(len(insertANode.Columns), 3, t)
	assert.Eq(insertANode.Columns[0], tutils.Table1AwesomeKeyId, t)
	assert.Eq(insertANode.Columns[1], tutils.Table1NonNullCollId, t)
	assert.Eq(insertANode.Columns[2], tutils.Table1NullableColId, t)

	//insertANode.Values
}
