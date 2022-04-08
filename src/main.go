package main

import (
	"fmt"
)

func main() {
	array := []int{0, 1, 2, 3, 4, 5}

	for _, integer := range array[1:] {
		print(integer)
	}

}

func PrintUsageInfo() {
	fmt.Println(
		"Database console ready.\n" +
			"Write command and sent it using ctrl+d")
}
