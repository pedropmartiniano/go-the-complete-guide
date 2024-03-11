package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	var investmentAmount float64 // deixando explicito o tipo da variavel
	expectedReturnRate := 5.5    // deixando implicito o tipo da variável, utilizando shortcut(:=)
	var years float64

	// & é um ponteiro para a variável "investmentAmount" que já foi declarada
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future value: %.2f", futureValue)
	outputText(formattedFV)

	fmt.Println("Future value adjusted for inflation: ", futureRealValue)
}

func outputText(text string) {
	fmt.Println(text)
}

// Ao declarar as variaveis no retorno da função(segundo parenteses), não é mais preciso instanciá-las dentro da função, nem especifica-las no retorno
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (futureValue, futureRealValue float64) {
	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)

	return futureValue, futureRealValue
	// somente "return" quando o nome das variaveis de retorno já estão declaradas
	// return
}
