class Solution:
    def findCircleNum(self, isConnected: List[List[int]]) -> int:
        res = 0
        visited = []

        def dfs(start: int) -> None:
            visited.append(start)
            for end in range(len(isConnected)):
                if isConnected[start][end] and end not in visited:
                    dfs(end)
        
        for i in range(len(isConnected)):
            if i not in visited:
                res += 1
                dfs(i)
        
        return res
