class Solution(object):
    def maxOperations(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        b = {}
        ans = 0
        for i in nums:
            if k - i in b and b[k - i] > 0:
                ans += 1
                b[k - i] -= 1
            elif i not in b:
                b[i] = 1
            else:
                b[i] += 1
        
        return ans