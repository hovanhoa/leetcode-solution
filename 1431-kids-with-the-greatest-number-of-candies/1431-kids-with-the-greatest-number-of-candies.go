func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandy := 0
    for _, v := range candies {
        maxCandy = max(maxCandy, v)
    }
    
    ans := make([]bool, len(candies))
    for i, v := range candies {
        if v + extraCandies >= maxCandy {
            ans[i] = true
        }
    }

    return ans
}

