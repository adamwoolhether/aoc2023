package day03

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
)

//go:embed schematic.txt
var schematic embed.FS

type schemaNum struct {
	row  int
	iLow int
	iTop int
	val  int
}

type symbolMap map[int][]int

func Run() (int, error) {
	data, err := schematic.Open("schematic.txt")
	if err != nil {
		return -1, fmt.Errorf("opening calibration document: %w", err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	potentialPartNumbers := make([]schemaNum, 0)
	symbols := make(map[int][]int)

	row := 0
	for scanner.Scan() {
		if scanner.Err() != nil {
			return -1, fmt.Errorf("scanning calibration document: %w", err)
		}

		line := scanner.Bytes()

		potentialPartNumbers = append(potentialPartNumbers, analyzeLine(line, row, symbols)...)

		row++
	}

	part1 := getPartNumbers(potentialPartNumbers, symbols)

	return part1, nil
}

func isTouching(symbolLoc, low, top int) bool {
	return (symbolLoc-1 >= low && symbolLoc-1 <= top) ||
		(symbolLoc+1 >= low && symbolLoc+1 <= top) ||
		symbolLoc == low || symbolLoc == top
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func isSymbol(b byte) bool {
	chars := []byte("#$+*%@&-/=")
	return bytes.Contains(chars, []byte{b})
}

func byteToInt(input []byte) int {
	var result int

	for i := range input {
		result = result*10 + int(input[i]-'0')
	}

	return result
}
