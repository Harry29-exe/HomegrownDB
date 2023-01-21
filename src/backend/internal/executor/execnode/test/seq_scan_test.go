package exenode

import (
	"HomegrownDB/backend/internal/executor"
	"HomegrownDB/backend/internal/testinfr"
	"HomegrownDB/dbsystem/tx"
	"HomegrownDB/hgtest"
	"testing"
)

func TestSeqScan_SimpleSelect(t *testing.T) {
	dbUtils := hgtest.CreateAndLoadDBWith(nil, t).WithUsersTable().Build()
	currentTx := dbUtils.DB.TxManager().New(tx.CommittedRead)

	inputQuery := "SELECT u.id, u.name FROM users u"
	plan := testinfr.ParseAnalyseAndPlan(inputQuery, dbUtils.DB.TableStore(), t)

	executor.Execute(plan, currentTx, dbUtils.DB)
}
