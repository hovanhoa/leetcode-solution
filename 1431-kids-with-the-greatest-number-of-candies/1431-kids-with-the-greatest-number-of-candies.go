func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandy := 0
    for _, candy := range candies {
        if candy > maxCandy {
            maxCandy = candy
        }
    }

    res := []bool{}

    for _, candy := range candies {
        if candy + extraCandies >= maxCandy {
            res = append(res, true)
        } else {
            res = append(res, false)
        }
    }

    return res
}