package main

import (
	"fmt"
	"os"

	"aoc2023/days/day01"
	"aoc2023/days/day02"
	"aoc2023/days/day03"
)

func main() {
	day01Pt1Result, day01Pt2Result, err := day01.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Day 01 Error: %v\n", err)
		os.Exit(1)
	}

	day02Pt1Result, day02Pt2Result, err := day02.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Day 02 Error: %v\n", err)
		os.Exit(1)
	}

	day03Res, err := day03.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Day 02 Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Day 01 Pt 1Result: %d\n", day01Pt1Result)
	fmt.Printf("Day 01 Pt 2 Result: %d\n", day01Pt2Result)
	fmt.Printf("Day 02 Pt 1 Result: %d\n", day02Pt1Result)
	fmt.Printf("Day 02 Pt 2 Result: %d\n", day02Pt2Result)
	fmt.Printf("Day 03 Pt 1 Result: %d\n", day03Res)

}
