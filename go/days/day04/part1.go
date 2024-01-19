package day04

import (
	"encoding/hex"
	"math"
)

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
