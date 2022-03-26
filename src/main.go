package main

import (
	"HomegrownDB/sql/schema/dbtable"
	"HomegrownDB/utils"
	"fmt"
)

func main() {
	s := dbtable.GetColumnType(dbtable.Int2, nil)
	s1 := *s

	s.IsFixedSize = false
	print(s.IsFixedSize)
	print((*s).IsFixedSize)
	print(s1.IsFixedSize)

	counter := utils.NewLockCounter(0)
	println("\n")
	val0 := counter.GetAndIncrement()
	println(val0)
	val2 := counter.IncrementAndGet()
	println(val2)
}

func PrintUsageInfo() {
	fmt.Println(
		"Database console ready.\n" +
			"Write command and sent it using ctrl+d")
}
