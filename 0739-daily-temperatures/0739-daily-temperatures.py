class Solution:
    def dailyTemperatures(self, temperatures: List[int]) -> List[int]:
        n = len(temperatures)
        res = [0] * n
        stack = []

        for i, val in enumerate(temperatures):
            while stack and val > temperatures[stack[-1]]:
                indx = stack.pop()
                res[indx] = i - indx
            stack.append(i)
        
        return res
