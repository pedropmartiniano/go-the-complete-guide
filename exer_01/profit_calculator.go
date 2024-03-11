package main

import (
	"fmt"
)

func main() {
	revenue := getUserInputs("Revenue: ")
	expenses := getUserInputs("Expenses: ")
	tax := getUserInputs("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, tax)

	fmt.Printf("Earnings before tax: %.2f\n", ebt)
	fmt.Printf("Earnings after tax: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func getUserInputs(infoText string) float64 {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

// return ebt, profit, ratio
func calculateFinancials(revenue float64, expenses float64, tax float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - (tax / 100 * ebt)
	ratio = ebt / profit

	return ebt, profit, ratio
}
