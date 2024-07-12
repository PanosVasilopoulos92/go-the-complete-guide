package main

import "fmt"

func main() {
	age := 32

	// var agePointer *int
	agePointer := &age

	// adultsYears := getAdultsYears(age)
	// fmt.Println(adultsYears)

	adultsYears := getAdultsYears(agePointer)
	fmt.Println(adultsYears)
	fmt.Println(age)
}

func getAdultsYears(age *int) int {
	*age = *age - 18 // will change the underlying (original) value
	return *age - 18
}
