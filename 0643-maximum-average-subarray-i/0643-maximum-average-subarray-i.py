class Solution(object):
    def findMaxAverage(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: float
        """
        maxSum = sum(nums[:k])
        curSum = maxSum

        for i in range(k, len(nums)):
            curSum += nums[i] - nums[i-k]
            print(curSum)
            maxSum = max(maxSum, curSum)

        return float(maxSum) / k