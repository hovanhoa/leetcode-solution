func replaceElements(arr []int) []int {
    ans := make([]int, len(arr))
    max := -1
    for i := len(arr)-1; i >= 0; i-- {
        ans[i] = max

        if max < arr[i] {
            max = arr[i]
        }
    }   

    return ans
}