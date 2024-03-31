class Solution(object):
    def removeStars(self, s):
        """
        :type s: str
        :rtype: str
        """
        res = []
        for c in s:
            if c != "*":
                res.append(c)
            else:
                if len(res) > 0:
                    res.pop()
        
        return "".join(res)
        