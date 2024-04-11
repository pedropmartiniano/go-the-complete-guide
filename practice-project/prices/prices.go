package prices

import (
	"fmt"

	"example.com/practice-project/conversion"
	"example.com/practice-project/manager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64           `json:"tax_rate"`
	Prices            []float64         `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
	Manager           manager.Manager   `json:"-"` // excluir esse campo do json criado a partir dessa estrutura
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.Manager.ReadLines()

	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		return err
	}

	job.Prices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process() error {
	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.Prices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	// essa função já irá retornar um erro ou nil
	return job.Manager.WriteResult(job)
}

func NewTaxIncludedPriceJob(manager manager.Manager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
		Manager: manager,
	}
}
