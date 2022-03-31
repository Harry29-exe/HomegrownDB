package main

import (
	"HomegrownDB/sql/schema/column/types"
	"fmt"
)

func main() {
	col := types.Int2Column{}
	serializedColumn := col.Serialize()

	print(serializedColumn)
}

func PrintUsageInfo() {
	fmt.Println(
		"Database console ready.\n" +
			"Write command and sent it using ctrl+d")
}
