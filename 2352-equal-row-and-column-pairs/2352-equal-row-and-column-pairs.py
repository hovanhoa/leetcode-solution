class Solution(object):
    def equalPairs(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        n = len(grid)
        hashMap = defaultdict(int)

        for row in grid:
            rowStr = str(row)
            hashMap[rowStr] += 1
        
        res = 0
        for j in range(n):
            col = [grid[i][j] for i in range(n)]
            colStr = str(col)
            res += hashMap.get(colStr, 0)

        return res