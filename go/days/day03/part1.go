package day03

func getPartNumbers(potentials []schemaNum, symbols symbolMap) int {
	var result int
	
outerLoop:
	for i := range potentials {
		pot := potentials[i]
		rows := []int{-1, 0, 1}

		for _, v := range rows {
			for _, symbolIdx := range symbols[pot.row+v] {
				if isTouching(symbolIdx, pot.iLow, pot.iTop) {
					result += pot.val
					continue outerLoop
				}
			}
		}
	}

	return result
}
