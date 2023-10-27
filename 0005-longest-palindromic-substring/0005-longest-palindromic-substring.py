class Solution:
    def longestPalindrome(self, s: str) -> str:
        n = len(s)
        maxLen = 1
        start = 0
        
        for i in range(n):
            low = i - 1
            high = i + 1

            while low >= 0 and s[low] == s[i]:
                low -= 1

            while high < n and s[high] == s[i]:
                high += 1

            while low >= 0 and high < n and s[high] == s[low]:
                low, high = low - 1, high + 1
            
            newLen = high - low - 1
            if newLen > maxLen:
                maxLen = newLen
                start = low + 1
        
        return s[start:start+maxLen]
