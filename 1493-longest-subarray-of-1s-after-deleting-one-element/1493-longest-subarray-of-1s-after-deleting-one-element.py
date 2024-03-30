class Solution(object):
    def longestSubarray(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        res = 0
        left = 0
        countZero = 0

        for right in range(len(nums)):
            if nums[right] == 0:
                countZero += 1
            
            while countZero == 2:
                if nums[left] == 0:
                    countZero -= 1
                left += 1

            res = max(res, right - left)

        return res