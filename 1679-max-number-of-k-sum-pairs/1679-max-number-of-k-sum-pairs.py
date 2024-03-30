class Solution(object):
    def maxOperations(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        numsMap = {}
        ans = 0
        for i in nums:
            if k - i in numsMap and numsMap[k - i] > 0:
                ans += 1
                numsMap[k - i] -= 1
            elif i not in numsMap:
                numsMap[i] = 1
            else:
                numsMap[i] += 1

        return ans