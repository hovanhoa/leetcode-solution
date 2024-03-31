class Solution(object):
    def asteroidCollision(self, asteroids):
        """
        :type asteroids: List[int]
        :rtype: List[int]
        """
        res = []
        for a in asteroids:
            while res and res[-1] > 0 > a:
                if res[-1] < abs(a):
                    res.pop()
                    continue
                elif res[-1] == abs(a):
                    res.pop()
                break
            else:
                res.append(a)
        
        return res
