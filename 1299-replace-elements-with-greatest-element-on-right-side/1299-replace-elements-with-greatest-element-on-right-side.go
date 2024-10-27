func replaceElements(arr []int) []int {
    ans := make([]int, len(arr))
    ans[len(arr)-1] = -1
    max := arr[len(arr)-1]
    for i := len(arr)-2; i >= 0; i-- {
        ans[i] = max

        if max < arr[i] {
            max = arr[i]
        }
    }   

    return ans
}