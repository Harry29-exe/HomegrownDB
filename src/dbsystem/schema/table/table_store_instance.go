package table

var DBTableStore Store

func init() {
	DBTableStore = NewEmptyTableStore()
	//todo init Store variable
}
