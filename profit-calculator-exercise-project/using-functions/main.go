package main

import (
	"fmt"
	"log"
)

func main() {
	revenue, expenses, taxRate := askUserToProvideInfo()
	ebt := calculateEBT(revenue, expenses)
	profit := calculateProfit(ebt, taxRate)
	ratio := calculateRatio(ebt, profit)

	fmt.Printf("Your EBT (earnings before tax) are: %.2f\n", ebt)
	fmt.Printf("Your profit (earnings after tax) are: %.2f\n", profit)
	fmt.Printf("Your ratio is: %.2f\n", ratio)

}

func getUserInput() float64 {
	var userInput float64

	_, err := fmt.Scan(&userInput)
	if err != nil {
		log.Fatalf("error while reading user input")
	}

	return float64(userInput)
}

func askUserToProvideInfo() (float64, float64, float64) {
	fmt.Print("What is your revenue? ")
	revenue := getUserInput()

	fmt.Print("What are your expenses? ")
	expenses := getUserInput()

	fmt.Print("What is your tax rate? ")
	taxRate := getUserInput()

	return revenue, expenses, taxRate
}

func calculateEBT(revenue float64, expenses float64) float64 {
	return revenue - expenses
}

func calculateProfit(ebt float64, taxRate float64) float64 {
	return ebt * (1 - taxRate/100)
}

func calculateRatio(ebt float64, profit float64) float64 {
	return ebt / profit
}
