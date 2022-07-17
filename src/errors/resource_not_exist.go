package errors

import "fmt"

type TableNotExist struct {
	TableName string
}

func (t TableNotExist) Error() string {
	return fmt.Sprintf("table with name: %s, does not exist.", t.TableName)
}
