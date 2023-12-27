package main

import (
	"fmt"
	"os"

	"aoc2023/days/day01"
	"aoc2023/days/day02"
)

func main() {
	day01Pt1Result, day01Pt2Result, err := day01.Sum()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Day 01 Error: %v\n", err)
		os.Exit(1)
	}

	day02Result, err := day02.Result()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Day 02 Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Day 01 Pt 1Result: %d\n", day01Pt1Result)
	fmt.Printf("Day 01 Pt 2 Result: %d\n", day01Pt2Result)
	fmt.Printf("Day 02 Result: %d\n", day02Result)
}
