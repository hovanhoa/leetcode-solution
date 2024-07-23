func kidsWithCandies(candies []int, extraCandies int) []bool {
    res := make([]bool, len(candies))
    maxCandy := 0

    for i := 0; i < len(candies); i += 1 {
        if candies[i] > maxCandy {
            maxCandy = candies[i]
        }
    }

    for i := 0; i < len(candies); i += 1 {
        if candies[i] + extraCandies >= maxCandy {
            res[i] = true
        }
    }

    return res
}