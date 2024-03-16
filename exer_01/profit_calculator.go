package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var revenue float64
	var expenses float64
	var tax float64
	var err error

	revenue, err = getUserInputs("Revenue: ")
	if err != nil {
		// or just
		// panic(err)
		for err != nil {
			revenue, err = getUserInputs("Revenue: ")
		}
	}

	expenses, err = getUserInputs("Expenses: ")
	if err != nil {
		// or just
		// panic(err)
		for err != nil {
			expenses, err = getUserInputs("Expenses: ")
		}
	}

	tax, err = getUserInputs("Tax Rate: ")
	if err != nil {
		// or just
		// panic(err)
		for err != nil {
			tax, err = getUserInputs("Tax Rate: ")
		}
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, tax)

	writeResultsIntoFile(ebt, profit, ratio)

	fmt.Printf("Earnings before tax: %.2f\n", ebt)
	fmt.Printf("Earnings after tax: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func getUserInputs(infoText string) (float64, error) {
	var userInput float64

	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0.0, errors.New("invalid value must be greater than 0")
		
	}

	return userInput, nil
}

// return ebt, profit, ratio
func calculateFinancials(revenue float64, expenses float64, tax float64) (ebt float64, profit float64, ratio float64) {
	ebt = revenue - expenses
	profit = ebt - (tax / 100 * ebt)
	ratio = ebt / profit

	return ebt, profit, ratio
}

func writeResultsIntoFile(ebt, profit, ratio float64) {
	fileName := "./results.txt"
	formatedResults := fmt.Sprintf("ebt=%.2f\nprofit=%.2f\nratio=%.2f\n", ebt, profit, ratio)
	os.WriteFile(fileName, []byte(formatedResults), 0644)
}