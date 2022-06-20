package main

import (
	"HomegrownDB/dbsystem/bstructs"
	"HomegrownDB/dbsystem/schema/column"
	"HomegrownDB/dbsystem/schema/column/ctypes"
	"HomegrownDB/dbsystem/schema/column/factory"
	"HomegrownDB/dbsystem/schema/table"
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
	tabDef := table.NewDefinition("table1", 231, 4352)
	tabDef.AddColumn(factory.CreateDefinition(column.NewSimpleArgs("col1", ctypes.Int2)))
	tabDef.AddColumn(factory.CreateDefinition(column.NewSimpleArgs("col2", ctypes.Int2)))
	tabDef.AddColumn(factory.CreateDefinition(
		column.ArgsBuilder().
			Name("col3").
			Type(ctypes.Int2).
			Nullable(true).
			Build()))
	tabDef.AddColumn(factory.CreateDefinition(
		column.ArgsBuilder().
			Name("col4").
			Type(ctypes.Int2).
			Nullable(true).
			Build()))

	colValues := map[string]any{
		"col1": 2,
		"col2": 4,
		"col3": 6,
		"col4": 9,
	}
	txCtx := tx.NewContext(231)
	tuple, err := bstructs.CreateTuple(tabDef, colValues, txCtx)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	strs := bstructs.TupleHelper.TupleDescription(tuple.Tuple)
	for _, str := range strs {
		println(str)
	}
}
