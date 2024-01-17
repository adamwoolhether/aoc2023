package main

import (
	"fmt"
	"os"

	"aoc2023/days/day01"
	"aoc2023/days/day02"
	"aoc2023/days/day03"
	"aoc2023/days/day04"
)

func main() {
	day01Pt1Result, day01Pt2Result, err := day01.Run()
	handleError(01, err)

	day02Pt1Result, day02Pt2Result, err := day02.Run()
	handleError(02, err)

	day03Pt1Res, day03Pt2Res, err := day03.Run()
	handleError(03, err)

	day4Res, err := day04.Run()
	handleError(04, err)

	fmt.Printf("Day 01 Pt 1 Result: %d\n", day01Pt1Result)
	fmt.Printf("Day 01 Pt 2 Result: %d\n", day01Pt2Result)
	fmt.Printf("Day 02 Pt 1 Result: %d\n", day02Pt1Result)
	fmt.Printf("Day 02 Pt 2 Result: %d\n", day02Pt2Result)
	fmt.Printf("Day 03 Pt 1 Result: %d\n", day03Pt1Res)
	fmt.Printf("Day 03 Pt 2 Result: %d\n", day03Pt2Res)
	fmt.Printf("Day 04 Pt 1 Result: %d\n", day4Res)

}

func handleError(dayNum int, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Day %d Error: %v\n", dayNum, err)
		os.Exit(1)
	}
}
