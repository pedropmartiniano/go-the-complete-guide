package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"1", "2"}
	productNames[3] = "4"

	prices := [4]float64{10.99, 9.99, 45.99, 20.00}

	// slices
	// slices são partes do array original, porém com os mesmos endereços de memória dos valores do array original, ou seja, ao modificar um indice de um slice, também estará modificando aquele indice no array original

	// pegar os elementos a partir do indice 1, até o indice 3(onde o indice 3 não está incluso)
	featuredPrices := prices[1:3]
	// pegar do inicio até o indice indicado(não incluido)
	featuredPrices2 := featuredPrices[:3]
	// pegar do indice indicado(incluso) até o final
	featuredPrices3 := featuredPrices2[2:]

	// todos os slices e o array original serão alterados
	featuredPrices3[0] = 79.90
	// featuredPrices3 = append(featuredPrices3, 123.21)

	// por mais que featuredPrices tenha apenas o indice 1 e 2 do array original, ele guarda uma cópia do array original consigo, e ao realizar um slice dele, é possivel acessar os outros elementos do array original que ele não possui, porém apenas os elementos a sua direita, pois os elementos da esquerda do array original que ele não possui não será possível acessar
	// aqui foi possível perceber que ele conseguiu recuperar o ultimo item do array original que ele não possuia, porém não foi possivel recuperar o primeiro item do array que ele não possuia
	featuredPrices = featuredPrices[:3]

	fmt.Println(featuredPrices)
	fmt.Println(featuredPrices2)
	fmt.Println(featuredPrices3)
	// capacidade será de 2 pois foi "cortado" os dois primeiros indices do array original
	fmt.Println(cap(featuredPrices3))
	fmt.Println(productNames) // indice 2 está vazio
	fmt.Println(prices)
	// tamanho do array e capacidade do array
	fmt.Println(len(prices), cap(prices))

	fmt.Println()
	// array de tamanho dinâmico
	teste := []string{"1", "2"}
	fmt.Println(teste)

	for i := 0; i < 10; i++ {
		teste = append(teste, fmt.Sprint(i))
	}

	

}
