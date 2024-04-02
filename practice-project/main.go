package main

import (
	"fmt"

	"example.com/practice-project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	// result := make(map[float64]prices.TaxIncludedPriceJob)

	for _, tax := range taxRates {
		pricesJob := prices.NewTaxIncludedPriceJob(tax)
		pricesJob.Process()
		fmt.Println(pricesJob)
	}

}
