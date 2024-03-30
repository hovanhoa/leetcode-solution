class Solution(object):
    def maxVowels(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: int
        """
        vowels = "aeiou"
        maxLen = 0
        curLen = 0

        for i in range(len(s)):
            if i < k and s[i] in vowels:
                curLen += 1
            
            if i >= k:
                if s[i-k] in vowels:
                    curLen -= 1
                
                if s[i] in vowels:
                    curLen += 1

            maxLen = max(maxLen, curLen)

        return maxLen
                