func canJump(nums []int) bool {
    cur := 0
    for i, v := range nums {
        if cur < i {
            return false
        }

        if i + v > cur {
            cur = i + v
        } 
    }

    return cur >= len(nums) - 1
}