package main

import (
	"fmt"
)

func main() {
	t := test{myMap: map[string]int{}}
	fmt.Println(t)

	t.Add("5", 5)
	fmt.Println(t)

	t.Rename("super name")
	fmt.Println(t)

	var tI1 ITest = test{myMap: map[string]int{}}
	tI2 := tI1.Rename("cool name")
	tI1.Add("4", 4)
	fmt.Println(tI1)
	fmt.Println(tI2)
	tI3 := tI2

	switch val := tI2.(type) {
	case test:
		val.name = "changed"
	}

	fmt.Println("---")
	fmt.Println(tI2)
	fmt.Println(tI3)
}

type ITest interface {
	Rename(str string) ITest
	Add(str string, i int)
}

type test struct {
	myMap map[string]int
	name  string
}

func (t test) Rename(str string) ITest {
	t.name = str
	return t
}

func (t test) Add(str string, i int) {
	t.myMap[str] = i
}

func PrintUsageInfo() {
	fmt.Println(
		"Database console ready.\n" +
			"Write command and sent it using ctrl+d")
}
