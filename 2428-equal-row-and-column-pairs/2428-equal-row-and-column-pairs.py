class Solution:
    def getRow(self, grid: List[List[int]], num: int) -> List[int]:
        res = []
        for i in grid:
            res.append(i[num])

        return res

    def equalPairs(self, grid: List[List[int]]) -> int:
        ans = 0

        for i in range(len(grid)):
            for j in range(len(grid)):        
                if grid[i] == self.getRow(grid, j):
                    ans += 1
                
        return ans
        