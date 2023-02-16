package analyse_test

import (
	"HomegrownDB/backend/internal/analyser"
	"HomegrownDB/backend/internal/node"
	"HomegrownDB/backend/internal/parser"
	. "HomegrownDB/backend/internal/testinfr"
	"HomegrownDB/dbsystem/hgtype"
	"HomegrownDB/dbsystem/reldef/tabdef"
	"HomegrownDB/dbsystem/reldef/tabdef/column"
	"HomegrownDB/hgtest"
	"HomegrownDB/lib/tests/assert"
	"testing"
)

var CreateTable = createTable{}

type createTable struct{}

// -------------------------
//      Positive
// -------------------------

func TestCreateTable_SimpleUsersTable(t *testing.T) {
	inputQuery := `CREATE TABLE users (
 		username VARCHAR(255),
 		surname VARCHAR(255) 
	);`
	db := hgtest.CreateAndLoadDBWith(nil, t).Build()
	pTree, err := parser.Parse(inputQuery)
	assert.ErrIsNil(err, t)

	query, err := analyser.Analyse(pTree, db.DB.AccessModule().RelationManager())
	assert.ErrIsNil(err, t)

	NodeAssert.Eq(
		CreateTable.usersTableATree(t),
		query,
		t,
	)
}

func (createTable) usersTableATree(t *testing.T) node.Query {
	query := node.NewQuery(node.CommandTypeUtils, nil)
	table := tabdef.NewDefinition("users")
	err := table.AddColumn(column.NewDefinition(
		"username", 0, 0, hgtype.NewStr(hgtype.Args{
			Length:   255,
			Nullable: true,
			VarLen:   true,
			UTF8:     true,
		})))
	assert.ErrIsNil(err, t)

	err = table.AddColumn(column.NewDefinition(
		"surname", 0, 0, hgtype.NewStr(hgtype.Args{
			Length:   255,
			Nullable: true,
			VarLen:   true,
			UTF8:     true,
		})))
	assert.ErrIsNil(err, t)

	query.UtilsStmt = node.NewCreateRelationTable(table)

	return query
}

func TestCreateTable_UsersTableWithAge(t *testing.T) {
	inputQuery := `CREATE TABLE users (
 		username VARCHAR(255),
 		surname VARCHAR(255),
		age INT8
	);`
	db := hgtest.CreateAndLoadDBWith(nil, t).Build()
	pTree, err := parser.Parse(inputQuery)
	assert.ErrIsNil(err, t)

	query, err := analyser.Analyse(pTree, db.DB.AccessModule().RelationManager())
	assert.ErrIsNil(err, t)

	NodeAssert.Eq(
		CreateTable.usersTableWithAgeATree(t),
		query,
		t,
	)
}

func (createTable) usersTableWithAgeATree(t *testing.T) node.Query {
	query := node.NewQuery(node.CommandTypeUtils, nil)
	table := tabdef.NewDefinition("users")
	err := table.AddColumn(column.NewDefinition(
		"username", 0, 0, hgtype.NewStr(hgtype.Args{
			Length:   255,
			Nullable: true,
			VarLen:   true,
			UTF8:     true,
		})))
	assert.ErrIsNil(err, t)

	err = table.AddColumn(column.NewDefinition(
		"surname", 0, 0, hgtype.NewStr(hgtype.Args{
			Length:   255,
			Nullable: true,
			VarLen:   true,
			UTF8:     true,
		})))
	assert.ErrIsNil(err, t)

	err = table.AddColumn(column.NewDefinition(
		"age", 0, 0, hgtype.NewInt8(hgtype.Args{
			Nullable: true,
		})))
	assert.ErrIsNil(err, t)

	query.UtilsStmt = node.NewCreateRelationTable(table)

	return query
}

// -------------------------
//      Negative
// -------------------------

func TestCreateTable_UsersTableWithAgeWrongType(t *testing.T) {
	inputQuery := `CREATE TABLE users (
 		username VARCHAR(255),
 		surname VARCHAR(255),
		age INTX
	);`
	db := hgtest.CreateAndLoadDBWith(nil, t).Build()
	pTree, err := parser.Parse(inputQuery)
	assert.ErrIsNil(err, t)

	_, err = analyser.Analyse(pTree, db.DB.AccessModule().RelationManager())
	assert.ErrNotNil(err, t)
}

func TestCreateTable_UsersTableWithAge_LengthOnINT8(t *testing.T) {
	inputQuery := `CREATE TABLE users (
 		username VARCHAR(255),
 		surname VARCHAR(255),
		age INT8(2)
	);`
	db := hgtest.CreateAndLoadDBWith(nil, t).Build()
	pTree, err := parser.Parse(inputQuery)
	assert.ErrIsNil(err, t)

	_, err = analyser.Analyse(pTree, db.DB.AccessModule().RelationManager())
	assert.ErrNotNil(err, t)
}
