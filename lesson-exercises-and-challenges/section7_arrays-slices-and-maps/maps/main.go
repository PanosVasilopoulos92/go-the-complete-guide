package main

import (
	"fmt"
	"strings"
)

func main() {
	websites := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["LinkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "Google")
	fmt.Println(websites)
	fmt.Println(strings.Repeat("-", 10))

	for k, v := range websites {
		fmt.Printf("%s -> %s\n", k, v)
	}
}
