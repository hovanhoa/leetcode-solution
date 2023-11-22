class Solution:
    def findDiagonalOrder(self, nums: List[List[int]]) -> List[int]:
        res = defaultdict(deque)
        for i, row in enumerate(nums):
            for j, x in enumerate(row):
                res[i + j].appendleft(x)
        return chain(*res.values())