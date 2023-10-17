class Solution:
    def countBits(self, n: int) -> List[int]:
        res = []
        for i in range(n + 1):
            res.append(sum([int(x) for x in bin(i)[2:]]))
        return res 
