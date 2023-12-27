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

func Run() (int, int, error) {
	var part1Result, part2Result int

	data, err := gamesInput.Open("games.txt")
	if err != nil {
		return -1, -1, fmt.Errorf("opening calibration document: %w", err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		if scanner.Err() != nil {
			return -1, -1, fmt.Errorf("scanning calibration document: %w", err)
		}

		lineBytes := scanner.Bytes()

		part1Result += possibleGames(lineBytes)
		part2Result += maxPowerGames(lineBytes)
	}

	return part1Result, part2Result, nil
}

func possibleGames(line []byte) int {
	gameID, rounds, _ := bytes.Cut(line, []byte(":"))

	results := bytes.Split(rounds, []byte(";"))

	for i := range results {
		cubes := bytes.Split(results[i], []byte(","))

		for j := range cubes {
			trimmed := bytes.TrimSpace(cubes[j])
			amount, color, _ := bytes.Cut(trimmed, []byte(" "))

			amt := byteToInt(amount)

			if amt > limits[string(color)] {
				return 0
			}
		}
	}

	successfulAmt, _ := bytes.CutPrefix(gameID, []byte("Game "))

	return byteToInt(successfulAmt)
}

func maxPowerGames(line []byte) int {
	_, rounds, _ := bytes.Cut(line, []byte(":"))

	results := bytes.Split(rounds, []byte(";"))

	highestCubeCount := make(map[string]int, 3)

	for i := range results {
		cubes := bytes.Split(results[i], []byte(","))

		for j := range cubes {
			trimmed := bytes.TrimSpace(cubes[j])
			amount, color, _ := bytes.Cut(trimmed, []byte(" "))
			colorStr := string(color)

			amt := byteToInt(amount)

			stored, ok := highestCubeCount[colorStr]
			if !ok || amt > stored {
				highestCubeCount[colorStr] = amt
			}
		}
	}

	result := 1
	for _, v := range highestCubeCount {
		result *= v
	}

	return result
}

func byteToInt(input []byte) int {
	var result int

	for i := range input {
		result = result*10 + int(input[i]-'0')
	}

	return result
}
