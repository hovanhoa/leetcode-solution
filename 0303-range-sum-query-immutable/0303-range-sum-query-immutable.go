type NumArray struct {
    nums []int
}


func Constructor(nums []int) NumArray {
    sumArr := make([]int, 1)
    for i, v := range nums {
        sumArr = append(sumArr, sumArr[i] + v)
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