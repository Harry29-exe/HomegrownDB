package table

var DBTableStore *TableStore

func init() {
	DBTableStore = NewEmptyTableStore()
	//todo init Store variable
}
