package main

import (
	"fmt"
	"thought-for-the-day/pkg/thoughts"
)

func main() {
	fmt.Println(thoughts.GetRandomQuote())
}
