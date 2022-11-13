package planer_test

import (
	"HomegrownDB/backend/internal/analyser"
	"HomegrownDB/backend/internal/analyser/anode"
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/backend/internal/shared/qctx"
	"HomegrownDB/common/tests/assert"
	"testing"
)

func TestBasicSelect1(t *testing.T) {
	query := "SELECT b.name, b.species FROM birds b"
	//ctx := qctx.NewQueryCtx()
	//ptree := BasicSelect1.parse(query, txCtx, t)
	//atree := BasicSelect1.analyse(ptree, txCtx, t)
	//todo implement me
	panic("Not implemented")
	println(query)
}

type basicSelect1 struct{}

var BasicSelect1 = basicSelect1{}

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
	_ = root
	//assert.Eq(1, len(root.Tables.Tables), t)
	//ptreeRoot := ptree.Root.(pnode.Select)
	//assert.Eq(ptreeRoot.Tables[0].TableName, root.Tables.Tables[0].Def.Name(), t)
	//assert.Eq(ptreeRoot.Tables[0].TableAlias, root.Tables.Tables[0].Alias, t)

	//todo implement me
	panic("Not implemented")
}
