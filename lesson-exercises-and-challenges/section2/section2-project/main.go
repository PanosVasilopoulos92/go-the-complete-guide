package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var accountBalance float64
var exitProgram bool

func init() {
	createLogger()
}

func main() {
	fmt.Println("Welcome to our e-banking app.")

	for !exitProgram {
		displayMenu()
		handleUserDecision()
	}

}

func displayMenu() {
	fmt.Println("What you wanna do?")
	fmt.Println("1. Show balance.")
	fmt.Println("2. Make withdrawal.")
	fmt.Println("3. Make deposit.")
	fmt.Println("4. Exit app.")
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
		writeTransactionToFile(logMsg)
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
		writeTransactionToFile(logMsg)
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

// createLogger creates a log file for transactions if it doesn't already exist.
// If the file doesn't exist, it creates a new file named "transactionsLogger.txt".
// If there is an error creating the file, it panics.
func createLogger() {
	_, err := os.Stat("transactionsLogger.txt")

	if os.IsNotExist(err) {
		_, err = os.Create("transactionsLogger.txt")
		if err != nil {
			panic(err)
		}
	}
}

// writeTransactionToFile appends the provided message to a file named "transactionsLogger.txt".
// If the file does not exist, it will be created. If there is an error opening or writing to the file,
// the function will log a fatal error.
func writeTransactionToFile(msg string) {
	file, err := os.OpenFile("transactionsLogger.txt", os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(msg + "\n")
	if err != nil {
		log.Fatal(err)
	}
}
