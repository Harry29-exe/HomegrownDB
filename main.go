package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	buffer := ""

	for {
		fmt.Print("db > ")
		buffer, _ = reader.ReadString('\004')
		fmt.Println(buffer)
	}
}

func PrintUsageInfo() {
	fmt.Println(
		"Database console ready.\n" +
			"Write command and sent it using ctrl+d")
}
