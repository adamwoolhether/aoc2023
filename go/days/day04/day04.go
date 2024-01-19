package day04

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
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
