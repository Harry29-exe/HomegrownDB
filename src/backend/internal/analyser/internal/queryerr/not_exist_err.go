package queryerr

import "fmt"

func ColumnNotExist(colName string, tableName string) NotExistError {
	return NotExistError{
		ObjType:   ObjTypeColumn,
		ObjName:   colName,
		ParentObj: tableName,
	}
}

func TableNotExistError(tableName string) NotExistError {
	return NotExistError{
		ObjType: ObjTypeTable,
		ObjName: tableName,
	}
}

type NotExistError struct {
	ObjType   string
	ObjName   string
	ParentObj string
}

func (n NotExistError) Error() string {
	if n.ParentObj == "" {
		return fmt.Sprintf("%s %s does not exist", n.ObjType, n.ObjName)
	}
	return fmt.Sprintf("%s %s does not exist on %s", n.ObjType, n.ObjName, n.ParentObj)
}

const (
	ObjTypeColumn = "Column"
	ObjTypeTable  = "Table"
)
