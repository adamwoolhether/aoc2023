package day05

func findLowestLocationPart2(seeds []int, mappings map[string][]Mapping) int {
	lowest := 0

	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < seeds[i]+seeds[i+1]; j++ {
			soil := getDest(j, mappings[seedToSoil])
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
	}

	return lowest
}
