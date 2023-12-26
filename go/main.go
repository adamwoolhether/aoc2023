package main

import (
	"fmt"
	"os"

	"aoc2023/days/day01"
)

func main() {
	day01Result, err := day01.Sum()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Day 01 Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Day 01 Result: %d\n", day01Result)
}
