class Solution(object):
    def kidsWithCandies(self, candies, extraCandies):
        """
        :type candies: List[int]
        :type extraCandies: int
        :rtype: List[bool]
        """
        greatestCandy = max(candies)
        res = []
        for candy in candies:
            if candy + extraCandies >= greatestCandy:
                res.append(True)
            else:
                res.append(False)
        
        return res