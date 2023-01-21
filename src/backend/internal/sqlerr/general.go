package sqlerr

import (
	"HomegrownDB/backend/internal/pnode"
	"fmt"
)

var _ error = NoTableWithNameErr{}

func NewNoTableWithNameErr(tableName string) NoTableWithNameErr {
	return NoTableWithNameErr{TableName: tableName}
}

type NoTableWithNameErr struct {
	TableName string
}

func (n NoTableWithNameErr) Error() string {
	return fmt.Sprintf("No table with name: %s", n.TableName)
}

func NewNoTableWithColErr(colName string) NoTableWithCol {
	return NoTableWithCol{ColName: colName}
}

type NoTableWithCol struct {
	ColName string
}

func (n NoTableWithCol) Error() string {
	return fmt.Sprintf("No table with column: %s", n.ColName)
}

// -------------------------
//      IllegalPNodeErr
// -------------------------

func NewIllegalPNodeErr(illegalNode pnode.Node, reason string) IllegalPNodeErr {
	return IllegalPNodeErr{
		Node:   illegalNode,
		Reason: reason,
	}
}

type IllegalPNodeErr struct {
	Node   pnode.Node
	Reason string
}

func (i IllegalPNodeErr) Error() string {
	return fmt.Sprintf("Illegal node found: %+v, reason: %s", i.Node, i.Reason)
}
