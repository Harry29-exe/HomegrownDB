package parser_test

import (
	"HomegrownDB/backend/internal/parser/parse"
	"HomegrownDB/backend/internal/parser/validator"
	"HomegrownDB/backend/internal/pnode"
	"HomegrownDB/lib/tests/assert"
	"testing"
)

var CreateTable = createTableUtils{}

type createTableUtils struct{}

func TestCreateTable_SimpleTable(t *testing.T) {
	query := `CREATE TABLE users (
 		username VARCHAR(255),
 		surname VARCHAR(255) );
	`
	src := newTestTokenSource(query)
	v := validator.NewValidator(src)

	parseTree, err := parse.CreateTable.Parse(src, v)
	assert.ErrIsNil(err, t)
	assert.Eq(len(src.Checkpoints), 0, t)

	assert.True(CreateTable.SimpleUsersTable().Equal(parseTree), t)
}

func (createTableUtils) SimpleUsersTable() pnode.CreateTableStmt {
	return pnode.NewCreateTableStmt("users", []pnode.ColumnDef{
		pnode.NewColumnDef("username", "VARCHAR", []pnode.TypeArg{
			pnode.NewArgLength(255),
		}),
		pnode.NewColumnDef("surname", "VARCHAR", []pnode.TypeArg{
			pnode.NewArgLength(255),
		}),
	})
}

func TestTableWithMultipleColumnDefinitions(t *testing.T) {
	query := `CREATE TABLE users ( 
		username VARCHAR(255), 
		surname VARCHAR(255), 
		age INT 
	);`
	src := newTestTokenSource(query)
	v := validator.NewValidator(src)

	parseTree, err := parse.CreateTable.Parse(src, v)
	assert.ErrIsNil(err, t)
	assert.Eq(len(src.Checkpoints), 0, t)

	assert.True(CreateTable.UsersTableWithAgeCol().Equal(parseTree), t)
}

func (createTableUtils) UsersTableWithAgeCol() pnode.CreateTableStmt {
	return pnode.NewCreateTableStmt("users", []pnode.ColumnDef{
		pnode.NewColumnDef("username", "VARCHAR", []pnode.TypeArg{
			pnode.NewArgLength(255),
		}),
		pnode.NewColumnDef("surname", "VARCHAR", []pnode.TypeArg{
			pnode.NewArgLength(255),
		}),
		pnode.NewColumnDef("age", "INT", []pnode.TypeArg{}),
	})
}

// -------------------------
//      negative
// -------------------------

func TestCreateTable_InvalidQuery(t *testing.T) {
	query := "CREATE TABLE"
	src := newTestTokenSource(query)
	v := validator.NewValidator(src)

	_, err := parse.CreateTable.Parse(src, v)
	assert.ErrNotNil(err, t)
}

func TestCreateTable_InvalidColumnName(t *testing.T) {
	query := "CREATE TABLE users ( 1username VARCHAR(255), surname VARCHAR(255) );"
	src := newTestTokenSource(query)
	v := validator.NewValidator(src)

	_, err := parse.CreateTable.Parse(src, v)
	assert.ErrNotNil(err, t)
}

func TestInvalidTypeArguments(t *testing.T) {
	query := "CREATE TABLE users ( username VARCHAR, surname VARCHAR(255, 10) );"
	src := newTestTokenSource(query)
	v := validator.NewValidator(src)

	_, err := parse.CreateTable.Parse(src, v)
	assert.ErrNotNil(err, t)
}

func TestMissingTableName(t *testing.T) {
	query := "CREATE TABLE ( username VARCHAR(255), surname VARCHAR(255) );"
	src := newTestTokenSource(query)
	v := validator.NewValidator(src)

	_, err := parse.CreateTable.Parse(src, v)
	assert.ErrNotNil(err, t)
}
