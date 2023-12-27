package day01

import (
	"bufio"
	"embed"
	"fmt"
)

//go:embed calibration.txt
var calibrationDoc embed.FS

func Sum() (int, int, error) {
	var part1Result, part2Result int

	data, err := calibrationDoc.Open("calibration.txt")
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

		part1Result += extractValsPart1(lineBytes)
		part2Result += extractValsPart2(lineBytes)
	}

	return part1Result, part2Result, nil
}

func isDigit(b byte) (bool, int) {
	var intDigit int

	isInt := b >= '0' && b <= '9'

	if isInt {
		intDigit = int(b - '0')
	}

	return isInt, intDigit
}
