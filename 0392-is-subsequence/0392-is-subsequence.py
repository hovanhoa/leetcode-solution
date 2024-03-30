class Solution(object):
    def isSubsequence(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        if s == "":
            return True
        
        indx = 0
        for c in t:
            if indx < len(s) and s[indx] == c:
                indx += 1
        
        return indx == len(s)