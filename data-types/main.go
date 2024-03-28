package main

import (
	"fmt"
)

// type alias
type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// make previne realocação de memória em slices e maps se usado da maneira correta

	// maneira de criar um array dinamico com um tamanho inicial, que pode ser ultrapassado
	// assim, golang cria um array de 2 posições vazias, e somente após ele instanciar a terceira posição, que ele é recriado como um slice, e a cada nova posição, é criado um novo array com aquela quantidade de posição
	// terceiro parametro define o capacidade maxima do array(opcional)
	userNames := make([]string, 2, 5)

	userNames[0] = "maxine"
	userNames[1] = "manuela"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	for index, value := range userNames {
		fmt.Println("Index", index, "value:", value)
	}

	// go irá evitar de realocar memória nos primeiros 3 registros do map, já que foi predefinido que ele irá ter pelo menos 10 atributos
	coursesRatings := make(floatMap, 3)

	coursesRatings["go"] = 4.8
	coursesRatings["react"] = 4.5
	coursesRatings["nodejs"] = 4.9

	coursesRatings.output()

	for key, value := range coursesRatings {
		fmt.Println("Index", key, "value:", value)
	}
}
