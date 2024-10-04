func predictPartyVictory(senate string) string {
    radiants, dires := []int{}, []int{}
    for i, v := range senate {
        if v == 'R' {
            radiants = append(radiants, i)
        } else {
            dires = append(dires, i)
        }
    }

    for len(radiants) > 0 && len(dires) > 0 {
        r, d := radiants[0], dires[0]
        if r < d {
            radiants = append(radiants, len(senate) + r)
        } else {
            dires = append(dires, len(senate) + d)
        }

        radiants = radiants[1:]
        dires = dires[1:]
    }

    if len(radiants) > 0 {
        return "Radiant"
    }


    return "Dire"
}