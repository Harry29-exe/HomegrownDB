package main

import (
	"HomegrownDB/sql/schema"
	"fmt"
)

func main() {
	s := schema.GetColumnType(schema.Int2, nil)
	s1 := *s

	s.IsFixedSize = false
	print(s.IsFixedSize)
	print((*s).IsFixedSize)
	print(s1.IsFixedSize)

}

func PrintUsageInfo() {
	fmt.Println(
		"Database console ready.\n" +
			"Write command and sent it using ctrl+d")
}
