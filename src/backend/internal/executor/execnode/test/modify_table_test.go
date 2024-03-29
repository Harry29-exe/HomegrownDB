package exenode_test

import (
	"HomegrownDB/backend/internal/executor"
	"HomegrownDB/backend/internal/testinfr"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/hgtest"
	"testing"
)

func TestModifyTable_SimpleInsert(t *testing.T) {
	dbUtils := hgtest.CreateAndLoadDBWith(nil, t).WithUsersTable().Build()
	currentTx := dbUtils.DB.TxManager().New(tx.CommittedRead)

	inputQuery := "INSERT INTO users (id, name) VALUES (1, 'bob')"
	plan := testinfr.ParseAnalyseAndPlan(inputQuery, dbUtils.DB.RelationManager(), t)

	executor.Execute(plan, currentTx, dbUtils.DB.ExecutionContainer())
}
