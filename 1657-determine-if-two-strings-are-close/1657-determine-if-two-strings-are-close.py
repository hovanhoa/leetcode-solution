class Solution(object):
    def closeStrings(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: bool
        """
        if len(word1) != len(word2):
            return False

        map1={}
        map2={}
        for i in word1:
            map1[i]=map1.get(i,0)+1
        for i in word2:
            map2[i]=map2.get(i,0)+1

        if set(map1.keys()) != set(map2.keys()):
            return False

        a = list(map1.values())
        b = list(map2.values())
        a.sort()
        b.sort()

        return a == b
        