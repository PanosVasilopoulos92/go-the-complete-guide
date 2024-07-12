package main

import (
	"fmt"
	"lesson-exercises-and-challenges/section3_packages/filesop"
	"time"
)

var accountBalance float64
var exitProgram bool

func init() {
	filesop.CreateLogger()
}

func main() {
	fmt.Println("Welcome to our e-banking app.")

	for !exitProgram {
		displayMenu()
		handleUserDecision()
	}

}

func getUserChoice() int {
	var userChoice int
	fmt.Scan(&userChoice)

	return userChoice
}

// handleUserDecision is the main function that handles the user's choice from the program's menu.
// It prompts the user for their choice, and then executes the corresponding action based on the user's input.
// The function updates the accountBalance variable and prints relevant information to the console.
func handleUserDecision() {
	var userChoice = getUserChoice()

	switch userChoice {
	case 1:
		fmt.Printf("Your account balance is : %.2f euros\n", accountBalance)
	case 2:
		makeWithdrawal()
	case 3:
		makeDeposit()
	case 4:
		fmt.Println("Ok, come back soon :)")
		exitProgram = true
	}
}

// makeDeposit is a function that handles the logic for depositing money into an account.
// It prompts the user for the amount to deposit, validates the amount, and updates the account balance
// if the deposit is successful.
func makeDeposit() {
	var amount float64

	fmt.Print("Please provide the amount of money you wish to deposit: ")
	fmt.Scan(&amount)
	fmt.Println()

	if isValidAmount(amount) {
		accountBalance += amount
		fmt.Printf("Transaction completed. %.2f euros were added to your account\n", amount)

		logMsg := fmt.Sprintf(" %v: %.2f euros were added.", time.Now().Format("2006-01-02 15:04:05"), amount)
		filesop.WriteTransactionToFile(logMsg)
	} else {
		fmt.Println("Sorry, not a valid amount. Try again.")
	}
}

// makeWithdrawal is a function that handles the logic for withdrawing money from an account.
// It prompts the user for the amount to withdraw, validates the amount, and updates the account balance
// if the withdrawal is successful.
func makeWithdrawal() {
	var amount float64

	fmt.Print("Please provide the amount of money you wish to withdraw: ")
	fmt.Scan(&amount)

	if !isValidAmount(amount) {
		fmt.Println("Sorry, not a valid amount. Try again.")
	} else if amount > accountBalance {
		fmt.Println("Sorry, account balance insufficient")
	} else {
		accountBalance -= amount
		fmt.Printf("Transaction completed. %.2f euros were removed from your account\n", amount)

		logMsg := fmt.Sprintf(" %v: %.2f euros were withdrew.", time.Now().Format("2006-01-02 15:04:05"), amount)
		filesop.WriteTransactionToFile(logMsg)
	}
}

// isValidAmount checks if the given amount is greater than 0.
// It returns true if the amount is valid, and false otherwise.
func isValidAmount(amount float64) bool {
	if amount <= 0 {
		return false
	} else {
		return true
	}
}
