class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        if not s:
            return True

        i = 0
        j = 0

        while j < len(t):
            if t[j] == s[i]:
                i += 1
            j += 1

            if i == len(s):
                return True
        
        return False
