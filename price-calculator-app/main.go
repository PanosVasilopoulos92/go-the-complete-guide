package main

import (
	"price-calculator-app/prices"
)

func main() {
	taxRates := []float64{0.0, 0.07, 0.1, 0.2}
	// results := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate)
		priceJob.Process()
	}

	// for k, v := range results {
	// 	fmt.Printf("For %.2f tax rate, prices are: %.02f respectively\n", k, v)
	// }
}
