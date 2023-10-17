class Solution:
    def tribonacci(self, n: int) -> int:
        res = []
        num = 0
        while num <= n:
            if num == 0:
                res.append(0)
            elif num in [1, 2]:
                res.append(1)
            else:
                res.append(sum(res[num-3: num]))
            num += 1

        return res[-1]
