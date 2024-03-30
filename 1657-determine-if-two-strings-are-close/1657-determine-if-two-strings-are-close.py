class Solution(object):
    def closeStrings(self, word1, word2):
        """
        :type word1: str
        :type word2: str
        :rtype: bool
        """
        if len(word1)!=len(word2):
            return False
        d1={}
        d2={}
        for i in word1:
            d1[i]=d1.get(i,0)+1
        for i in word2:
            d2[i]=d2.get(i,0)+1
        if set(d1.keys())!=set(d2.keys()):
            return False
        a=list(d1.values())
        b=list(d2.values())
        a.sort()
        b.sort()
        return a==b
        
        