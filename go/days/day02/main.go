package day02

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
)

var limits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

//go:embed games.txt
var gamesInput embed.FS

func Result() (int, error) {
	var result int

	data, err := gamesInput.Open("games.txt")
	if err != nil {
		return -1, fmt.Errorf("opening calibration document: %w", err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		if scanner.Err() != nil {
			return -1, fmt.Errorf("scanning calibration document: %w", err)
		}

		lineBytes := scanner.Bytes()

		result += possibleGames(lineBytes)
	}

	return result, nil
}

func possibleGames(line []byte) int {
	before, after, _ := bytes.Cut(line, []byte(":"))

	results := bytes.Split(after, []byte(";"))

	for i := range results {
		picks := bytes.Split(results[i], []byte(","))

		for j := range picks {
			trimmed := bytes.TrimSpace(picks[j])
			amount, color, _ := bytes.Cut(trimmed, []byte(" "))

			amt := byteToInt(amount)

			if amt > limits[string(color)] {
				return 0
			}
		}
	}

	successfulAmt, _ := bytes.CutPrefix(before, []byte("Game "))

	return byteToInt(successfulAmt)
}

func byteToInt(input []byte) int {
	var result int

	for i := range input {
		result = result*10 + int(input[i]-'0')
	}

	return result
}
