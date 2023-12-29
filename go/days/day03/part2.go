package day03

type gearLoc struct {
	row int
	idx int
}

func getGearRatios(potentials []schemaNum, gears symbolMap) int {
	var result int

	matchGears := make(map[gearLoc][]int)

outerLoop:
	for i := range potentials {
		pot := potentials[i]
		rows := []int{-1, 0, 1}

		for _, v := range rows {
			for _, gearIdx := range gears[pot.row+v] {

				if isTouching(gearIdx, pot.iLow, pot.iTop) {
					gl := gearLoc{row: pot.row + v, idx: gearIdx}
					matchGears[gl] = append(matchGears[gl], pot.val)
					continue outerLoop
				}
			}
		}
	}

	for _, v := range matchGears {

		if len(v) == 2 {
			//fmt.Println(v, v[0], v[1])
			result += v[0] * v[1]
		}
	}

	return result
}
