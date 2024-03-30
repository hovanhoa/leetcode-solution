class Solution(object):
    def closeStrings(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: bool
        """
        if len(word1) != len(word2):
            return False
        
        l = [word1.count(i) for i in set(word1)]
        m = [word2.count(i) for i in set(word2)]
        return set(word1) == set(word2) and sorted(l) == sorted(m)