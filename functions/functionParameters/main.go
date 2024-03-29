package functionparameters

import (
	"fmt"
)

// alias para tipagem
type transformFn func(int) int

func main() {
	numbers1 := []int{1, 2, 3, 4}
	numbers2 := []int{5, 6, 7, 8}

	// passando a funcão como parametro
	doubleNum := transformNumbers(&numbers1, double)
	tripleNum := transformNumbers(&numbers1, triple)

	fmt.Println(doubleNum)
	fmt.Println(tripleNum)

	fmt.Println("----------------------------")

	// getTransformerFunction irá retornar qual função ele irá usar(double ou triple) de acordo com o primeiro número do array, depois irei passar o retorno dessa função(que é uma das duas funções citadas) como parametro para a função trasformNumbers
	transformerFn1 := getTransformerFunction(&numbers1)
	transformerFn2 := getTransformerFunction(&numbers2)
	differentNumbers1 := transformNumbers(&numbers1, transformerFn1)
	differentNumbers2 := transformNumbers(&numbers2, transformerFn2)

	fmt.Println(differentNumbers1)
	fmt.Println(differentNumbers2)
}

// recebendo funções como parâmetros
// o parametro função é uma funcao que recebe um tipo inteiro e retorna um tipo inteiro
func transformNumbers(numbers *[]int, transform transformFn) []int {
	var dNumber []int

	for _, val := range *numbers {
		dNumber = append(dNumber, transform(val))
	}

	return dNumber
}

// retornando uma função como valor
func getTransformerFunction(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(num int) int {
	return num * 2
}

func triple(num int) int {
	return num * 3
}
