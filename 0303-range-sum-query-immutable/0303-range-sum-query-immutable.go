type NumArray struct {
    nums []int
}


func Constructor(nums []int) NumArray {
    return NumArray{
        nums: nums,
    }
}


func (this *NumArray) SumRange(left int, right int) int {
    sum := 0
    for left <= right && right < len(this.nums) {
        sum += this.nums[left]
        left += 1
    }

    return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */