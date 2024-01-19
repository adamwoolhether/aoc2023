package day04

func winAmt(winningAmt []int) int {
	var total int

	totalWins := make([]int, len(winningAmt))
	for i := range totalWins {
		totalWins[i] = 1
	}

	for i, amt := range winningAmt {
		for j := 1; j <= amt && i+j < len(winningAmt); j++ {
			totalWins[i+j] += totalWins[i]
		}
	}

	for _, count := range totalWins {
		total += count
	}

	return total
}
