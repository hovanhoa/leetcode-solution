class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        res = 0
        currEnd = float("-inf")

        for i in sorted(intervals, key=lambda x: x[1]):
            if i[0] >= currEnd:
                currEnd = i[1]
            else:
                res += 1
        
        return res
