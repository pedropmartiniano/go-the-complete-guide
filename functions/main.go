package main

import "fmt"

func main() {
	num := factorial(5)

	fmt.Println(num)
}

func factorial(num int) int {
	if num == 1 {
		return 1
	}

	return num * factorial(num-1)
}
