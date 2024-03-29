package closure

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// criando funções a partir do closure
	double := createTransformer(2)
	triple := createTransformer(3)

	// passando uma função anônima como parâmetro
	transformed := transformNumbers(&numbers, func(num int) int {
		return num * 2
	})

	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println(transformed)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// closure, uma função que produz outras funções com diferentes "configurações"
// a função retornada é capaz de armazenar o valor de "factor" de seu escopo superior, graças ao conceito de closure
func createTransformer(factor int) func(int) int {
	// função que multiplica o numero por x
	return func(num int) int {
		return num * factor
	}
}
