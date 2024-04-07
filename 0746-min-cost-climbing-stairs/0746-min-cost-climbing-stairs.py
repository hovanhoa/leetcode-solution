class Solution(object):
    def minCostClimbingStairs(self, cost):
        """
        :type cost: List[int]
        :rtype: int
        """
        cost.append(0)

        for i in range(2, len(cost)):
            cost[i] += min(cost[i-1], cost[i-2])
        
        return cost[-1]
        