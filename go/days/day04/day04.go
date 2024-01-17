package day04

import (
	"bufio"
	"bytes"
	"embed"
	"encoding/hex"
	"fmt"
	"math"
)

//go:embed scratchcards.txt
var scratchcards embed.FS

func Run() (int, int, error) {
	scratchers, err := scratchcards.Open("scratchcards.txt")
	if err != nil {
		return 0, 0, fmt.Errorf("opening scratchcards document: %w", err)
	}
	defer scratchers.Close()

	scanner := bufio.NewScanner(scratchers)

	var part1 int
	var part2 int

	winCardsAmt := make([]int, 0)

	cardsProcesses := 0
	for ; scanner.Scan(); cardsProcesses++ {
		if scanner.Err() != nil {
			return 0, 0, fmt.Errorf("reading scratchcards document: %w", scanner.Err())
		}

		line := scanner.Bytes()

		firstHalf, secondHalf, found := bytes.Cut(line, []byte("|"))
		if !found {
			return 0, 0, fmt.Errorf("could not find separator")
		}

		_, winningHalf, found := bytes.Cut(firstHalf, []byte(":"))
		if !found {
			return 0, 0, fmt.Errorf("could not find separator")
		}

		winningNumBytes := bytes.Fields(winningHalf)
		myNumBytes := bytes.Fields(secondHalf)

		matches, score := cardScore(winningNumBytes, myNumBytes)

		winCardsAmt = append(winCardsAmt, matches)

		part1 += score
	}

	part2 = winAmt(winCardsAmt)

	return part1, part2, nil
}

func cardScore(winningNumBytes [][]byte, myNumBytes [][]byte) (matches int, score int) {
	winningNum := make(map[string]struct{})

	for _, val := range winningNumBytes {
		winningNum[hex.EncodeToString(val)] = struct{}{}
	}

	for _, val := range myNumBytes {
		if _, found := winningNum[hex.EncodeToString(val)]; found {
			matches++
		}
	}

	//score := int(math.Pow(2, float64(matches)))- 1 // This adds them all up, not what we want here.
	score = int(math.Pow(2, float64(matches-1)))

	return matches, score
}

func winAmt(winningAmt []int) int {
	var total int

	totalWins := make([]int, len(winningAmt))
	for i := range totalWins {
		totalWins[i] = 1
	}

	for i, amt := range winningAmt {
		for j := 1; j <= amt && i+j < len(winningAmt); j++ {
			totalWins[i+j] += totalWins[i]
		}
	}

	for _, count := range totalWins {
		total += count
	}

	return total
}
