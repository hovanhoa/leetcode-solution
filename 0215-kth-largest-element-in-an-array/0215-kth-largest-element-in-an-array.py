class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        pivot = nums[len(nums)//2]
        l = [i for i in nums if i > pivot]
        m = [i  for i in nums if i == pivot]
        h = [i for i in nums if i < pivot]
        if k <= len(l):
            return self.findKthLargest(l,k)
        elif k > (len(l)+len(m)):
            return self.findKthLargest(h,k-(len(l)+len(m)))
        else:
            return m[0]