package main

import "fmt"

func main() {
	var first, second, third string
	first, second, third = "one", "two", "three"

	var fourth, fifth, sixth string = "four", "five", "six"

	seventh, eight, ninth := "seven", "eight", "nine"

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	fmt.Println(first, second, third, fourth, fifth, sixth, seventh, eight, ninth)
	fmt.Println(one, isFriday, twoPointTwo, say)
}
