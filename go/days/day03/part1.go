package day03

func analyzeLine(line []byte, row int, symbols symbolMap) []schemaNum {
	var parts []schemaNum

	for i := 0; i < len(line); i++ {
		if isSymbol(line[i]) {
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

func getPartNumbers(potentials []schemaNum, symbols symbolMap) int {
	var result int
outerLoop:
	for i := range potentials {
		pot := potentials[i]
		rows := []int{-1, 0, 1}

		for _, v := range rows {
			for _, v := range symbols[pot.row+v] {
				if isTouching(v, pot.iLow, pot.iTop) {
					result += pot.val
					continue outerLoop
				}
			}
		}
	}

	return result
}
