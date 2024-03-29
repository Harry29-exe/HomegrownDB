package main

import (
	"HomegrownDB/dbsystem/reldef"
	"HomegrownDB/dbsystem/reldef/tabdef/column"
	"HomegrownDB/dbsystem/reldef/tabdef/column/ctypes"
	"HomegrownDB/dbsystem/reldef/tabdef/column/factory"
	"HomegrownDB/dbsystem/storage/tpage"
	"HomegrownDB/dbsystem/tx"
	"os"
)

func main() {
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//page.Register(r)

	println("\n")
	testTuplePrint()
	println("\n")

	//err := r.Run()
	//if err != nil {
	//	os.Exit(1)
	//} // listen and serve on 0.0.0.0:8080
}

func testTuplePrint() {
	tableDef := reldef.CreateTableDefinition("ttable1")
	tableDef.SetOID(231)
	tableDef.SetOID(4352)

	tableDef.AddNewColumn(factory.CreateDefinition(column.ArgsBuilder("col1", ctypes.Int2).Build()))
	tableDef.AddNewColumn(factory.CreateDefinition(column.ArgsBuilder("col2", ctypes.Int2).Build()))
	tableDef.AddNewColumn(factory.CreateDefinition(
		column.ArgsBuilder("col3", ctypes.Int2).
			Nullable(true).
			Build()))
	tableDef.AddNewColumn(factory.CreateDefinition(
		column.ArgsBuilder("col4", ctypes.Int2).
			Nullable(true).
			Build()))

	colValues := map[string]any{
		"col1": 2,
		"col2": 4,
		"col3": 6,
		"col4": 9,
	}
	txCtx := tx.NewContext(32)
	tuple, err := tpage.NewTestTuple(tableDef, colValues, txCtx)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	strs := tpage.TupleDebugger.TupleDescription(tuple.Tuple)
	for _, str := range strs {
		println(str)
	}
}
