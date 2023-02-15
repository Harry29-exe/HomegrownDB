package analyse_test

import (
	"HomegrownDB/backend/internal/parser"
	"HomegrownDB/lib/tests/assert"
	"testing"
)

func TestCreateTable_SimpleTable(t *testing.T) {
	query := `CREATE TABLE users (
 		username VARCHAR(255),
 		surname VARCHAR(255) );
	`
	pTree, err := parser.Parse(query)
	assert.ErrIsNil(err, t)

	_ = pTree
	//analyser.Analyse(pTree)
}
