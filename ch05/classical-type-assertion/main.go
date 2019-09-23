package main

import (
	"fmt"
)

func HandleData(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("x is int")
	case string:
		fmt.Println("x is string")
	case map[string]int:
		fmt.Println("x is map[string]int")
	default:
		fmt.Println("x is unknown")
	}

}

func main() {
	HandleData(1)
	HandleData("1")
	HandleData(map[string]int{"foo": 1})
	HandleData(1.1)
}
