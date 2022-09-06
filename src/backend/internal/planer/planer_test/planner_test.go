package planer_test

import (
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/common/tests/assert"
	"testing"
)

func TestBasicSelectParse(t *testing.T) {
	query := "SELECT b.name, b.species FROM birds b"
	tree, err := parser.Parse(query)
	if err != nil {
		t.Errorf("Could not parse query becouse of error %s", err)
	}

	assert.Eq(parser.Select, tree.RootType, t)
	root, ok := tree.Root.(*pnode.SelectNode)
	assert.Eq(ok, true, t)

	fields := root.Fields.Fields
	assert.Eq(len(fields), 2, t)

	nameField := fields[0]
	assert.Eq(nameField.FieldName, "name", t)
	assert.Eq(nameField.FieldAlias, "name", t)
	assert.Eq(nameField.TableAlias, "b", t)

	speciesField := fields[1]
	assert.Eq(speciesField.FieldName, "species", t)
	assert.Eq(speciesField.FieldAlias, "species", t)
	assert.Eq(speciesField.TableAlias, "b", t)

	tables := root.Tables.Tables
	assert.Eq(len(tables), 1, t)
	birdsTable := tables[0]
	assert.Eq(birdsTable.TableName, "birds", t)
	assert.Eq(birdsTable.TableAlias, "b", t)
}

func TestBasicInsertParse(t *testing.T) {

}
