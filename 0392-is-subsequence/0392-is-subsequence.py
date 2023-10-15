class Solution(object):
    def isSubsequence(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        if not s:
            return True

        indx = 0
        for i in t:
            if i == s[indx]:
                indx += 1
            if indx == len(s):
                return True
        return False
        