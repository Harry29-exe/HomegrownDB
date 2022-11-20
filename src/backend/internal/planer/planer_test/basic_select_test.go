package planer_test

import (
	"HomegrownDB/backend/internal/analyser"
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/backend/internal/planer/planer_test/internal"
	"HomegrownDB/backend/internal/shared/qctx"
	"HomegrownDB/common/tests/assert"
	"HomegrownDB/common/tests/tutils/testtable/ttable2"
	"HomegrownDB/dbsystem/schema/table"
	"testing"
)

func TestBasicSelect1(t *testing.T) {
	query := "SELECT b.name, b.species FROM birds b"
	table2 := ttable2.Def(t)
	tableStore, _ := table.NewTableStore([]table.Definition{table2})
	ctx := qctx.NewQueryCtx(tableStore)

	helper := basicSelect1{aNode: internal.ANode{T: t}}
	ptree := helper.parse(query, ctx, t)
	atree := helper.analyse(ptree, ctx, t)

	_ = atree
}

type basicSelect1 struct {
	aNode internal.ANode
}

func (b basicSelect1) parse(query string, ctx qctx.QueryCtx, t *testing.T) parser.Tree {
	ptree, err := parser.Parse(query, ctx)
	if err != nil {
		t.Errorf("Could not parse query becouse of error %s", err)
	}

	assert.Eq(parser.Select, ptree.RootType, t)
	root, ok := ptree.Root.(pnode.Select)
	assert.Eq(ok, true, t)

	fields := root.Fields
	assert.Eq(len(fields), 2, t)

	nameField := fields[0]
	assert.Eq(nameField.FieldName, "name", t)
	assert.Eq(nameField.FieldAlias, "name", t)
	assert.Eq(nameField.TableAlias, "b", t)

	speciesField := fields[1]
	assert.Eq(speciesField.FieldName, "species", t)
	assert.Eq(speciesField.FieldAlias, "species", t)
	assert.Eq(speciesField.TableAlias, "b", t)

	tables := root.Tables
	assert.Eq(len(tables), 1, t)
	birdsTable := tables[0]
	assert.Eq(birdsTable.TableName, "birds", t)
	assert.Eq(birdsTable.TableAlias, "b", t)
	return ptree
}

func (b basicSelect1) analyse(ptree parser.Tree, ctx qctx.QueryCtx, t *testing.T) analyser.Tree {
	atree, err := analyser.Analyse(ptree, ctx)
	assert.IsNil(err, t)
	assert.Eq(atree.RootType, analyser.RootTypeSelect, t)
	root, ok := atree.Root.(anode.Select)
	assert.Eq(true, ok, t)

	pTreeRoot := ptree.Root.(pnode.Select)
	b.aNode.CmpTables(pTreeRoot.Tables, root.Tables, ctx)

	//todo implement me
	panic("Not implemented")
}
