package sqlerr

import "fmt"

var _ error = NoTableWithName{}

func NewNoTableWithName(tableName string) NoTableWithName {
	return NoTableWithName{TableName: tableName}
}

type NoTableWithName struct {
	TableName string
}

func (n NoTableWithName) Error() string {
	return fmt.Sprintf("No table with name: %s", n.TableName)
}
