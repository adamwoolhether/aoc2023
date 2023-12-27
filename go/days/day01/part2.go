package day01

import (
	"bytes"
)

func extractValsPart2(lineBytes []byte) int {
	var (
		foundLower, foundUpper bool
		lower, upper           int
	)

	headStart, tailEnd := 0, len(lineBytes)
	headEnd, tailStart := 0, len(lineBytes)-1

	for headStart <= tailStart {
		if foundLower && foundUpper {
			break
		}

		if !foundLower {
			foundDigit, n := isDigit(lineBytes[headEnd])
			if foundDigit {
				lower = n
				foundLower = true
			}

			foundStrDigit, n := isStringDigit(lineBytes[headStart:headEnd])
			if foundStrDigit {
				lower = n
				foundLower = true
			}
			headEnd++
		}

		if !foundUpper {
			foundDigit, n := isDigit(lineBytes[tailStart])
			if foundDigit {
				upper = n
				foundUpper = true
			}

			foundStrDigit, n := isStringDigit(lineBytes[tailStart:tailEnd])
			if foundStrDigit {
				upper = n
				foundUpper = true
			}
			tailStart--
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

func isStringDigit(b []byte) (bool, int) {
	switch {
	case bytes.Contains(b, []byte("one")):
		return true, 1
	case bytes.Contains(b, []byte("two")):
		return true, 2
	case bytes.Contains(b, []byte("three")):
		return true, 3
	case bytes.Contains(b, []byte("four")):
		return true, 4
	case bytes.Contains(b, []byte("five")):
		return true, 5
	case bytes.Contains(b, []byte("six")):
		return true, 6
	case bytes.Contains(b, []byte("seven")):
		return true, 7
	case bytes.Contains(b, []byte("eight")):
		return true, 8
	case bytes.Contains(b, []byte("nine")):
		return true, 9
	default:
		return false, -1
	}
}
