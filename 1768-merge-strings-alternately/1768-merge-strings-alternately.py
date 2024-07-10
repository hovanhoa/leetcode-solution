class Solution(object):
    def mergeAlternately(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: str
        """
        i, j = 0, 0
        res = ""

        while i < len(word1) or j < len(word2):
            if i < len(word1):
                res += word1[i]
                i += 1

            if j < len(word2):
                res += word2[j]
                j += 1
        
        return res
        