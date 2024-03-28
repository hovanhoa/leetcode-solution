class Solution(object):
    def reverseVowels(self, s):
        """
        :type s: str
        :rtype: str
        """
        vowels = "aeiouAEIOU"
        i = 0
        j = len(s)-1
        lst = list(s)

        while i < j:
            while vowels.find(lst[i]) == -1 and i < j:
                i+=1
            
            while vowels.find(lst[j]) == -1 and i < j:
                j-=1
            
            lst[i], lst[j] = lst[j], lst[i]
            
            i+=1
            j-=1
        
        return ''.join(lst)