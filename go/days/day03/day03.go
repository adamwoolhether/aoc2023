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

func Run() (int, int, error) {
	data, err := schematic.Open("schematic.txt")
	if err != nil {
		return -1, -1, fmt.Errorf("opening calibration document: %w", err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)

	potentialPartNumbers := make([]schemaNum, 0)
	symbols := make(map[int][]int)
	gears := make(map[int][]int)

	row := 0
	for scanner.Scan() {
		if scanner.Err() != nil {
			return -1, -1, fmt.Errorf("scanning calibration document: %w", err)
		}

		line := scanner.Bytes()

		potentialPartNumbers = append(potentialPartNumbers, analyzeLine(line, row, symbols, gears)...)

		row++
	}

	part1 := getPartNumbers(potentialPartNumbers, symbols)
	part2 := getGearRatios(potentialPartNumbers, gears)

	return part1, part2, nil
}

func analyzeLine(line []byte, row int, symbols, gears symbolMap) []schemaNum {
	var parts []schemaNum

	for i := 0; i < len(line); i++ {
		if isSymbol(line[i]) {
			if isGear(line[i]) {
				gears[row] = append(gears[row], i)
			}

			symbols[row] = append(symbols[row], i)

			continue
		}

		if isDigit(line[i]) {
			for j := i + 1; ; {
				if j >= len(line) {
					parts = append(parts, schemaNum{
						row:  row,
						iLow: i,
						iTop: j - 1,
						val:  byteToInt(line[i:]),
					})

					break
				}
				if !isDigit(line[j]) {
					parts = append(parts, schemaNum{
						row:  row,
						iLow: i,
						iTop: j - 1,
						val:  byteToInt(line[i:j]),
					})

					if isSymbol(line[j]) {
						symbols[row] = append(symbols[row], j)

						if isGear(line[j]) {
							gears[row] = append(gears[row], j)
						}
					}

					i = j
					break
				}

				j++
			}
		}
	}

	return parts
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

func isGear(b byte) bool {
	gear := byte('*')
	return gear == b
}

func byteToInt(input []byte) int {
	var result int

	for i := range input {
		result = result*10 + int(input[i]-'0')
	}

	return result
}
