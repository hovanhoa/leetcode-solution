func kidsWithCandies(candies []int, extraCandies int) []bool {
    result := make([]bool, len(candies))

    maxVal := candies[0]
    for _, val := range candies {
        if val > maxVal {
            maxVal = val
        }
    }

    for index, val := range candies {
        if val + extraCandies >= maxVal {
            result[index] = true
        }
    }

    return result
}