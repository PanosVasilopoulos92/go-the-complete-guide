package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	path := "C:\\Users\\viato\\Desktop\\Udemy Courses\\Golang\\Go - The Complete Guide\\lesson-exercises-and-challenges\\section2\\46_writing-to-files\\balance.txt"
	getAccountBalance(path)
}

func getAccountBalance(file string) {
	data, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	dataToString := string(data)

	dataToFloat, err := strconv.ParseFloat(dataToString, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dataToFloat)
}
