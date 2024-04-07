class Solution(object):
    def eraseOverlapIntervals(self, intervals):
        """
        :type intervals: List[List[int]]
        :rtype: int
        """
        res = 0
        curEnd = float("-inf")

        for i in sorted(intervals, key=lambda x: x[1]):
            if i[0] >= curEnd:
                curEnd = i[1]
            else:
                res += 1

        return res
        