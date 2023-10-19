class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        n = len(nums)
        rev = sum(nums[:k])
        res = rev

        for i in range(k , n):
            rev -= nums[i-k]
            rev += nums[i]
            res = max(res, rev)
        
        return res/k
