func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandy := 0
    for _, candy := range candies {
        if candy > maxCandy {
            maxCandy = candy
        }
    }

    res := make([]bool, len(candies))

    for i, candy := range candies {
        if candy + extraCandies >= maxCandy {
            res[i] = true
        }
    }

    return res
}