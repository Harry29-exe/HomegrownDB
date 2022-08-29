package planer_test

import (
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/backend/internal/parser/pnode"
	"HomegrownDB/common/tests"
	"testing"
)

func TestBasicSelectParse(t *testing.T) {
	query := "SELECT b.name, b.species FROM birds b"
	tree, err := parser.Parse(query)
	if err != nil {
		t.Errorf("Could not parse query becouse of error %s", err)
	}

	tests.AssertEq(parser.Select, tree.RootType, t)
	root, ok := tree.Root.(*pnode.SelectNode)
	tests.AssertEq(ok, true, t)

	fields := root.Fields.Fields
	tests.AssertEq(len(fields), 2, t)

	nameField := fields[0]
	tests.AssertEq(nameField.FieldName, "name", t)
	tests.AssertEq(nameField.FieldAlias, "name", t)
	tests.AssertEq(nameField.TableAlias, "b", t)

	speciesField := fields[1]
	tests.AssertEq(speciesField.FieldName, "species", t)
	tests.AssertEq(speciesField.FieldAlias, "species", t)
	tests.AssertEq(speciesField.TableAlias, "b", t)

	tables := root.Tables.Tables
	tests.AssertEq(len(tables), 1, t)
	birdsTable := tables[0]
	tests.AssertEq(birdsTable.TableName, "birds", t)
	tests.AssertEq(birdsTable.TableAlias, "b", t)
}
