type NumArray struct {
    nums []int
}


func Constructor(nums []int) NumArray {
    s := 0
    arr := make([]int, len(nums) + 1)
    arr[0] = 0
    for i, v := range nums {
        s += v
        arr[i+1] = s
    }
    
    return NumArray {
        nums: arr,
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    return this.nums[right+1] - this.nums[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */