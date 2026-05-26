type NumArray struct {
    nums []int
}


func Constructor(nums []int) NumArray {
    return NumArray{
        nums: nums,
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    sumArr := make([]int, len(this.nums)+1)
    s := 0
    for i, v := range this.nums {
        s += v
        sumArr[i+1] = s
    }

    return sumArr[right+1] - sumArr[left]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */