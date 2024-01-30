package day05

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"strconv"
	"strings"
)

type Mapping struct {
	src      int
	dest     int
	rangeAmt int
}

const (
	seedToSoil        = "seed-to-soil"
	soilToFertilizer  = "soil-to-fertilizer"
	fertilizerToWater = "fertilizer-to-water"
	waterToLight      = "water-to-light"
	lightToTemp       = "light-to-temperature"
	tempToHumidity    = "temperature-to-humidity"
	humidityToLoc     = "humidity-to-location"
)

//go:embed almanac.txt
var almanac embed.FS

func Run() (int, int, error) {
	input, err := almanac.Open("almanac.txt")
	if err != nil {
		return 0, 0, fmt.Errorf("opening almanac: %w", err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(scanTwoLines)
	scanner.Scan()

	seeds := make([]int, 4)
	_, seedInts, found := strings.Cut(scanner.Text(), "seeds: ")
	if !found {
		panic("seeds not found")
	}

	seeds = getInts(seedInts)

	mappings := make(map[string][]Mapping)

	for scanner.Scan() {
		mapLines := strings.Split(scanner.Text(), "\n")
		mapName := strings.Fields(mapLines[0])[0]

		for i := 1; i < len(mapLines); i++ {
			mappings[mapName] = append(mappings[mapName], getMapping(mapLines[i]))
		}
	}

	lowestLoc := findLowestLocation(seeds, mappings)
	lowestLocPart2 := findLowestLocationPart2(seeds, mappings)

	return lowestLoc, lowestLocPart2, nil
}

func getInts(s string) []int {
	ints := make([]int, 0)

	for _, intStr := range strings.Fields(s) {
		val, err := strconv.Atoi(intStr)
		if err != nil {
			panic(err)
		}

		ints = append(ints, val)
	}

	return ints
}

func getMapping(s string) Mapping {
	ints := getInts(s)
	if len(ints) != 3 {
		panic("mapping should have 3 values")
	}

	m := Mapping{
		src:      ints[1], // Note the input is 'dest, src, range'
		dest:     ints[0],
		rangeAmt: ints[2],
	}

	return m
}

// getDest is the key mapping logic for this puzzle.
func getDest(src int, mappings []Mapping) int {
	for _, m := range mappings {
		if m.src <= src && src < m.src+m.rangeAmt {
			return m.dest + (src - m.src)
		}
	}

	return src
}

func scanTwoLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	dropCR := func(data []byte) []byte {
		if len(data) > 0 && data[len(data)-1] == '\r' {
			return data[0 : len(data)-1]
		}
		return data
	}

	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		// We have a full newline-terminated line.
		return i + 2, dropCR(data[0:i]), nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), dropCR(data), nil
	}
	// Request more data.
	return 0, nil, nil
}
