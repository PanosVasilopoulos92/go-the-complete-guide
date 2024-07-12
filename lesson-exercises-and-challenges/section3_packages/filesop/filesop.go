package filesop

import (
	"log"
	"os"
)

// createLogger creates a log file for transactions if it doesn't already exist.
// If the file doesn't exist, it creates a new file named "transactionsLogger.txt".
// If there is an error creating the file, it panics.
func CreateLogger() {
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
func WriteTransactionToFile(msg string) {
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
