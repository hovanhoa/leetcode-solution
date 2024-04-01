class Solution(object):
    def tribonacci(self, n):
        """
        :type n: int
        :rtype: int
        """
        dp = [0] * (n + 1)

        for i in range(1, n + 1):
            if i < 3:
                dp[i] = 1
            else:
                dp[i] = sum(dp[i-3:i])
            
        return dp[-1]
        