package day01

import (
	"bufio"
	"embed"
	"fmt"
)

//go:embed calibration.txt
var calibrationDoc embed.FS

func Sum() (int, error) {
	var result int

	data, err := calibrationDoc.Open("calibration.txt")
	if err != nil {
		return -1, fmt.Errorf("opening calibration document: %w", err)
	}

	scanner := bufio.NewScanner(data)
	//scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		if scanner.Err() != nil {
			return -1, fmt.Errorf("scanning calibration document: %w", err)
		}

		lineBytes := scanner.Bytes()

		result += extractValue(lineBytes)
	}

	return result, nil
}

func extractValue(lineBytes []byte) int {
	var (
		foundLower, foundUpper bool
		lower, upper           byte
	)

	start, end := 0, len(lineBytes)-1

	for start <= end {
		if foundLower && foundUpper {
			break
		}

		if !foundLower {
			if isDigit(lineBytes[start]) {
				lower = lineBytes[start] - '0'
				foundLower = true
			}
			start++
		}

		if !foundUpper {
			if isDigit(lineBytes[end]) {
				upper = lineBytes[end] - '0'
				foundUpper = true
			}
			end--
		}
	}

	if !foundUpper {
		return int(lower)*10 + int(lower)
	}

	if !foundLower {
		return int(upper)*10 + int(upper)
	}

	return int(lower)*10 + int(upper)
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
