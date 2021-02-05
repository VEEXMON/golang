package main

import "fmt"

func main() {
	var point = 1

	if point == 10 {
		fmt.Println("passed with a perfect score")
	} else if point > 5 {
		fmt.Println("passed")
	} else if point == 4 {
		fmt.Println("almost pass")
	} else {
		fmt.Printf("failure. your score %d\n", point)
	}
}
