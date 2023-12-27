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

	potentialPartNumbers := make([]schemaNum, 0) // x, y0, y1 : part#
	symbols := make(map[int][]int)

	xAxis := 0
	for scanner.Scan() {
		if scanner.Err() != nil {
			return -1, fmt.Errorf("scanning calibration document: %w", err)
		}

		line := scanner.Bytes()

		analyze(line, xAxis, &potentialPartNumbers, symbols)

		xAxis++
	}

	var result int
outerLoop:
	for i := range potentialPartNumbers {
		pot := potentialPartNumbers[i]
		above := pot.row - 1
		level := pot.row
		below := pot.row + 1

		for _, v := range symbols[above] {
			if isTouching(v, pot.iLow, pot.iTop) {
				result += pot.val
				continue outerLoop
			}
		}

		for _, v := range symbols[level] {
			if isTouching(v, pot.iLow, pot.iTop) {
				result += pot.val
				continue outerLoop
			}
		}

		for _, v := range symbols[below] {
			if isTouching(v, pot.iLow, pot.iTop) {
				result += pot.val
				continue outerLoop
			}
		}
	}

	return result, nil
}

func isTouching(symbolLoc, low, top int) bool {
	return (symbolLoc-1 >= low && symbolLoc-1 <= top) ||
		(symbolLoc+1 >= low && symbolLoc+1 <= top) ||
		symbolLoc == low || symbolLoc == top
}

func analyze(line []byte, row int, parts *[]schemaNum, symbols symbolMap) {
	//i := 0
	for i := 0; i < len(line); i++ {
		if isSymbol(line[i]) {
			symbols[row] = append(symbols[row], i)

			continue
		}

		if isDigit(line[i]) {
			for j := i + 1; ; {
				if j >= len(line) {
					*parts = append(*parts, schemaNum{
						row:  row,
						iLow: i,
						iTop: j - 1,
						val:  byteToInt(line[i:]),
					})
					break
				}
				if !isDigit(line[j]) {
					*parts = append(*parts, schemaNum{
						row:  row,
						iLow: i,
						iTop: j - 1,
						val:  byteToInt(line[i:j]),
					})

					if isSymbol(line[j]) {
						symbols[row] = append(symbols[row], j)
					}

					i = j
					break
				}

				j++
			}
		}
	}

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
