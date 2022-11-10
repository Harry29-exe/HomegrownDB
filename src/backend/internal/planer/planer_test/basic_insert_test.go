package planer

import (
	"HomegrownDB/dbsystem/tx"
	"testing"
)

func TestBasicInsertParse(t *testing.T) {
	query := "INSERT INTO users (id, name) VALUES (1, 'bob'), (2, 'alice')"
	txCtx := tx.NewContext(52, nil)

}
