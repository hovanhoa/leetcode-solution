func kidsWithCandies(candies []int, extraCandies int) []bool {
    res := make([]bool, len(candies))
    maxCandy := 0

    for _, val := range candies {
        if val > maxCandy {
            maxCandy = val
        }
    }

    for idx, val := range candies {
        if val + extraCandies >= maxCandy {
            res[idx] = true
        }
    }

    return res
}