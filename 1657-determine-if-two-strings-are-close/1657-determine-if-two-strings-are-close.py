class Solution(object):
    def closeStrings(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: bool
        """
        if len(word1) != len(word2):
            return False

        map1 = collections.Counter(word1)
        map2 = collections.Counter(word2)

        if set(map1.keys()) != set(map2.keys()):
            return False

        a = list(map1.values())
        b = list(map2.values())
        a.sort()
        b.sort()

        return a == b
        