package main

import (
	"fmt"

	"com.example/price_calculator/filemanager"
	"com.example/price_calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.20, .10, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		priceJob.Process()
	}

}
