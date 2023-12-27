package day01

func extractValsPart1(lineBytes []byte) int {
	var (
		foundLower, foundUpper bool
		lower, upper           int
	)

	start, end := 0, len(lineBytes)-1

	for start <= end {
		if foundLower && foundUpper {
			break
		}

		if !foundLower {
			foundDigit, n := isDigit(lineBytes[start])
			if foundDigit {
				lower = n
				foundLower = true
			}
			start++
		}

		if !foundUpper {
			foundDigit, n := isDigit(lineBytes[end])
			if foundDigit {
				upper = n
				foundUpper = true
			}
			end--
		}
	}

	if !foundLower {
		return upper*10 + upper
	}

	if !foundUpper {
		return lower*10 + lower
	}

	return lower*10 + upper
}
