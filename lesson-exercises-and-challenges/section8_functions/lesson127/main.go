package main

/* Returning functions as values */

import (
	"fmt"
	"strings"
)

func main() {
	originalSlice1 := []int{1, 2, 3, 4, 6}
	fmt.Println(originalSlice1)

	doubledSlice := transformNumbers(&originalSlice1, double)
	fmt.Println(doubledSlice)

	tripledSlice := transformNumbers(&originalSlice1, triple)
	fmt.Println(tripledSlice)
	fmt.Println(strings.Repeat("-", 10))

	originalSlice2 := []int{8, 22, 23, 14, 46}
	fmt.Println(originalSlice2)

	transformerFn1 := getTransformerFunction(&originalSlice1)
	transformerFn2 := getTransformerFunction(&originalSlice2)

	transformedNumbers := transformNumbers(&originalSlice2, transformerFn1)
	transformedNumbers2 := transformNumbers(&originalSlice2, transformerFn2)
	fmt.Println(transformedNumbers)
	fmt.Println(transformedNumbers2)
}

func getTransformerFunction(numbers *[]int) func(int) int {

	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func transformNumbers(s *[]int, transform func(int) int) []int {
	tNumbers := make([]int, 0, 10)

	for _, v := range *s {
		tNumbers = append(tNumbers, transform(v))
	}

	return tNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
