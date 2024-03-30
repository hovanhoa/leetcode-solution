class Solution(object):
    def largestAltitude(self, gain):
        """
        :type gain: List[int]
        :rtype: int
        """
        res = 0
        cur = 0
        for i in gain:
            cur += i
            res = max(res, cur)

        return res
