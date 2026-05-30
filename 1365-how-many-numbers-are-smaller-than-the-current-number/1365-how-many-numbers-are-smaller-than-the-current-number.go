func smallerNumbersThanCurrent(nums []int) []int {
    // Step 1: Create a frequency array for numbers 0 to 100
    count := make([]int, 101)
    for _, num := range nums {
        count[num]++
    }

    // Step 2: Transform counts into running prefix sums
    // count[i] will store how many numbers are less than or equal to i
    for i := 1; i <= 100; i++ {
        count[i] += count[i-1]
    }

    // Step 3: Map the original values to their prefix sum answers
    result := make([]int, len(nums))
    for i, num := range nums {
        if num == 0 {
            result[i] = 0
        } else {
            // Numbers smaller than num are stored at count[num-1]
            result[i] = count[num-1]
        }
    }

    return result
}