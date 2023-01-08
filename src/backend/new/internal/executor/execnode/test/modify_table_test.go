package exenode_test

import (
	"HomegrownDB/backend/new/internal/executor/exinfr"
	"HomegrownDB/backend/new/internal/testinfr"
	"testing"
)

func TestModifyTable_SimpleInsert(t *testing.T) {
	inputQuery := "INSERT INTO users (id, name) VALUES (1, 'bob')"
	plan := testinfr.ParseAnalyseAndPlan(inputQuery, store, t)

	exCtx := exinfr.NewExCtx(
		plan)
	//todo implement me
	panic("Not implemented")
}
