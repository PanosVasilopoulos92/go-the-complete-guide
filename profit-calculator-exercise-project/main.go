package main

import "fmt"

// main is the entry point of the profit calculator program. It prompts the user for revenue, expenses, and tax rate,
// then calculates and prints the earnings before and after tax, as well as the ratio between them.
func main() {
	var revenue float64
	var expenses float64
	var taxRate float64
	var earningsBeforeTax float64
	var earningsAfterTax float64
	var ratio float64

	fmt.Print("What is your revenue? ")
	fmt.Scan(&revenue)

	fmt.Print("What are your expenses? ")
	fmt.Scan(&expenses)

	fmt.Print("What is your tax rate? ")
	fmt.Scan(&taxRate)

	earningsBeforeTax = revenue - expenses
	fmt.Println("Your earnings before tax are", earningsBeforeTax)

	earningsAfterTax = earningsBeforeTax * (1 - taxRate/100)
	fmt.Println("Your earnings before tax are", earningsAfterTax)

	ratio = earningsBeforeTax / earningsAfterTax
	// fmt.Println("So, your ratio is", ratio)
	fmt.Printf("So, your ratio is %.2f", ratio)
}
