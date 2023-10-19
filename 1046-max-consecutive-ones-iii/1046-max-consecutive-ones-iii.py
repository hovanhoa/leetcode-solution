class Solution:
    def longestOnes(self, nums: List[int], k: int) -> int:
        res = 0
        l = 0

        for r, val in enumerate(nums):
            if val == 0:
                k -= 1
            while k < 0:
                if nums[l] == 0:
                    k += 1
                l += 1
            res = max(res, r - l + 1)
        
        return res
