package main

import (
	"fmt"

	"example.com/practice-project/cmdmanager"
	// "example.com/practice-project/filemanager"
	"example.com/practice-project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15, 0.2}
	taxes := make(map[float64]prices.TaxIncludedPriceJob)

	for _, tax := range taxRates {
		// fileManager := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", 100*tax))
		// pricesJob2 := prices.NewTaxIncludedPriceJob(fileManager, tax)
		// pricesJob2.Process()
		cmd := cmdmanager.New()
		pricesJob := prices.NewTaxIncludedPriceJob(cmd, tax)
		err := pricesJob.Process()

		if err != nil {
			fmt.Println("Could not process job")
			fmt.Println(err)
			continue
		}

		taxes[tax] = *pricesJob
	}

	fmt.Println(taxes)

}
