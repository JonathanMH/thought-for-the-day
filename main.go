package main

import (
	"fmt"
	"thought-for-the-day/pkg/thoughts"
)

func main() {
	fmt.Printf("+++ %s +++ \n", thoughts.GetRandomQuote())
}
