package prices

import (
	"fmt"

	"com.example/price_calculator/conversion"
	"com.example/price_calculator/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := filemanager.ReadLines("prices.txt")

	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}

	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		fmt.Println("Error converting price to float: ", err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string, len(job.InputPrices))

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.3f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result
	filemanager.WriteJson(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{100, 20, 30},
	}
}
