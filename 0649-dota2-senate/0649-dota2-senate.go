func predictPartyVictory(senate string) string {
	radiants, dires := []int{}, []int{}
	for i, c := range senate {
		if c == 'R' {
			radiants = append(radiants, i)
		} else {
			dires = append(dires, i)
		}
	}

	for len(radiants) > 0 && len(dires) > 0 {
		r, d := radiants[0], dires[0]
		radiants, dires = radiants[1:], dires[1:]
		if r < d {
			radiants = append(radiants, r+len(senate))
		} else {
			dires = append(dires, d+len(senate))
		}
	}

	if len(radiants) > 0 {
		return "Radiant"
	}

	return "Dire"
}