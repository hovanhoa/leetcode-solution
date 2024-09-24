class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        n = len(temperatures)
        res = [0] * n
        stack = []

        for i, v in enumerate(temperatures):
            while stack and v > temperatures[stack[-1]]:
                temp = stack.pop()
                res[temp] = i - temp  
            stack.append(i)
        
        return res
