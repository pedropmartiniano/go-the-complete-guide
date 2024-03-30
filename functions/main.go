package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// spread operator
	// desestrurando o array e passando cada indice como um parametro
	sum := sumup(0, numbers...)
	sum2 := sumup(0, 1, 2, 3, 4, 5, 6, 7)

	fmt.Println(sum)
	fmt.Println(sum2)
}

// função que recebe N parametros
// irá colocar todos os parametro dentro de uma slice(array) de inteiro
func sumup(startingValue int, numbers ...int) int {
	for _, val := range numbers {
		startingValue += val
	}

	return startingValue
}
