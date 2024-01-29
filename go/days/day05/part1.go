package day05

func findLowestLocation(seeds []int, mappings map[string][]Mapping) int {
	lowest := 0

	for _, seed := range seeds {
		soil := getDest(seed, mappings[seedToSoil])
		fertilizer := getDest(soil, mappings[soilToFertilizer])
		water := getDest(fertilizer, mappings[fertilizerToWater])
		light := getDest(water, mappings[waterToLight])
		temp := getDest(light, mappings[lightToTemp])
		humidity := getDest(temp, mappings[tempToHumidity])
		location := getDest(humidity, mappings[humidityToLoc])

		if lowest == 0 {
			lowest = location
			continue
		}

		if location < lowest {
			lowest = location
		}
	}

	return lowest
}

// getDest is the key mapping logic for this puzzle.
func getDest(src int, mappings []Mapping) int {
	for _, m := range mappings {
		if m.src <= src && src <= m.src+m.rangeAmt {
			return m.dest + (src - m.src)
		}
	}

	return src
}
