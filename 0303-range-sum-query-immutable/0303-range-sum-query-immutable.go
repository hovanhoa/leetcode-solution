type NumArray struct {
    nums []int
}


func Constructor(nums []int) NumArray {
    sumArr := make([]int, len(nums)+1)
    s := 0
    for i, v := range nums {
        s += v
        sumArr[i+1] = s
    }

    return NumArray{
        nums: sumArr,
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