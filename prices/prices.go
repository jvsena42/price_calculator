package prices

import (
	"bufio"
	"fmt"
	"os"

	"com.example/price_calculator/conversion"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() {
	file, error := os.Open("prices.txt")

	if error != nil {
		fmt.Println("Could not open file: ", error)
		return
	}

	scanner := bufio.NewScanner(file)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err := scanner.Err()

	if err != nil {
		fmt.Println("Error reading file: ", err)
		file.Close()
		return
	}

	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		fmt.Println("Error converting price to float: ", err)
		file.Close()
		return
	}

	job.InputPrices = prices
	file.Close()
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string, len(job.InputPrices))

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.3f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{100, 20, 30},
	}
}
