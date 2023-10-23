class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        # Helper function to check if a cell is within the grid
        def check(i, j):
            if i < 0 or j < 0 or i >= m or j >= n:
                return False
            return True
        
        # Recursive function with memoization
        @lru_cache(maxsize=None)
        def helper(i, j):
            if not check(i, j):
                return 0
            else:
                if j == n - 1 and i == m - 1:
                    return 1
                else:
                    # Move right (i, j+1) and move down (i+1, j)
                    return helper(i + 1, j) + helper(i, j + 1)

        # Start from the top-left corner
        return helper(0, 0)