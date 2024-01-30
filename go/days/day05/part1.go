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
