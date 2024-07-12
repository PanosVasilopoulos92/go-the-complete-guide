package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 3.5
	var annualReturnRate float64 = 5.5
	// var investmentAmount float64 = 1000
	var investmentAmount float64
	// years := 10
	var years int

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years of the investment: ")
	fmt.Scan(&years)

	// var futureValue = float64(investmentAmount) * math.Pow(1+annualReturnRate/100, float64(years))
	var futureValue = investmentAmount * math.Pow(1+annualReturnRate/100, float64(years))
	fmt.Println(futureValue)
	// fmt.Println(investmentAmount)
	// fmt.Printf("%T", investmentAmount)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, float64(years))
	fmt.Println(futureRealValue)
}
