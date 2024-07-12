package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func saveAccountBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)

	err := os.WriteFile("balance.txt", []byte(balanceText), 0644)
	if err != nil {
		err = errors.New("error while writing to file")
		log.Fatalln(err)
	}

}

func main() {
	var accountBalance float64

	accountBalance += 3400
	accountBalance -= 200
	accountBalance += 600

	saveAccountBalanceToFile(accountBalance)
}
