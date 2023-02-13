package exenode

import (
	"HomegrownDB/backend/internal/executor"
	"HomegrownDB/backend/internal/testinfr"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/hgtest"
	"HomegrownDB/lib/bparse"
	"HomegrownDB/lib/tests/assert"
	"testing"
)

func TestSeqScan_SimpleSelect(t *testing.T) {
	dbUtils := hgtest.CreateAndLoadDBWith(nil, t).WithUsersTable().Build()

	insertTx := dbUtils.DB.TxManager().New(tx.CommittedRead)
	insertQuery := "INSERT INTO users (id, name) VALUES (1, 'Bob')"
	insertPlan := testinfr.ParseAnalyseAndPlan(insertQuery, dbUtils.DB.RelationManager(), t)
	insertResult := executor.Execute(insertPlan, insertTx, dbUtils.DB.ExecutionContainer())
	assert.Eq(1, len(insertResult), t)

	selectTx := dbUtils.DB.TxManager().New(tx.CommittedRead)
	selectQuery := "SELECT u.id, u.name FROM users u"
	selectPlan := testinfr.ParseAnalyseAndPlan(selectQuery, dbUtils.DB.RelationManager(), t)
	selectResult := executor.Execute(selectPlan, selectTx, dbUtils.DB.ExecutionContainer())
	assert.Eq(1, len(selectResult), t)

	resultRow := selectResult[0]
	// todo create utils for validating binary values
	assert.EqArray(bparse.Serialize.Int8(1), resultRow.ColValue(0), t) // users.id
	assert.EqArray([]byte("Bob"), resultRow.ColValue(1)[4:], t)        // users.name
}
