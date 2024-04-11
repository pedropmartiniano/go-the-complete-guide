package main

import (
	"fmt"

	// "example.com/practice-project/cmdmanager"
	"example.com/practice-project/filemanager"
	"example.com/practice-project/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15, 0.2}
	taxes := make(map[float64]prices.TaxIncludedPriceJob)
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for i, tax := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)

		fileManager := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", 100*tax))
		pricesJob := prices.NewTaxIncludedPriceJob(fileManager, tax)
		go pricesJob.Process(doneChans[i], errorChans[i])

		// cmd := cmdmanager.New()
		// pricesJob := prices.NewTaxIncludedPriceJob(cmd, tax)
		// err := pricesJob.Process()

		// if err != nil {
		// 	fmt.Println("Could not process job")
		// 	fmt.Println(err)
		// 	continue
		// }

		// irá receber uma struct vazia pois a keyword "go" não espera o valor retornar
		taxes[tax] = *pricesJob
	}

	// irá esperar todos os canais terem um retorno para continuarem após esse for loop
	// estou iterando sobre taxRates pois sei que o tamanho do array de chanels é do tamanho de taxRates, ou seja, irei passar por todos os chanels
	for i := range taxRates {
		//estrutura de controle para trabalhar com "chanels"
		select {
		// o primeiro valor que for retornado para o chanel de um dos cases, já é o suficiente para ele passar para o proximo chanel sem ficar esperando uma resposta, já que o outro case já retornou uma
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[i]:
			fmt.Println("Done!")
		}

	}

}
