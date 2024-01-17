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

func Run() (int, error) {
	scratchers, err := scratchcards.Open("scratchcards.txt")
	if err != nil {
		return 0, fmt.Errorf("opening scratchcards document: %w", err)
	}
	defer scratchers.Close()

	scanner := bufio.NewScanner(scratchers)

	var totalScore int
	for scanner.Scan() {
		if scanner.Err() != nil {
			return 0, fmt.Errorf("reading scratchcards document: %w", scanner.Err())
		}

		line := scanner.Bytes()

		firstHalf, secondHalf, found := bytes.Cut(line, []byte("|"))
		if !found {
			return 0, fmt.Errorf("could not find separator")
		}

		_, winningHalf, found := bytes.Cut(firstHalf, []byte(":"))
		if !found {
			return 0, fmt.Errorf("could not find separator")
		}

		winningNumBytes := bytes.Fields(winningHalf)
		myNumBytes := bytes.Fields(secondHalf)

		totalScore += cardScore(winningNumBytes, myNumBytes)
	}

	return totalScore, nil
}

func cardScore(winningNumBytes [][]byte, myNumBytes [][]byte) int {
	var matches int

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
	score := int(math.Pow(2, float64(matches-1)))

	return score
}
