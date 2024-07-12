package main

/* Functions as Values and Function Types*/

import "fmt"

type transformFn func(int) int

func main() {
	originalSlice := []int{1, 2, 3, 4, 6}
	fmt.Println(originalSlice)

	doubledSlice := transformNumbers(&originalSlice, double)
	fmt.Println(doubledSlice)

	tripledSlice := transformNumbers(&originalSlice, triple)
	fmt.Println(tripledSlice)
}

func transformNumbers(s *[]int, transform transformFn) []int {
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
