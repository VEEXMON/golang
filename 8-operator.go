package main

import "fmt"

func main() {
	var value = (((6+2)%3)*4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Printf("result %d (%t) \n", value, isEqual)
}
