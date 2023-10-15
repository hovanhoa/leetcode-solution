class Solution(object):
    def reverseVowels(self, s):
        """
        :type s: str
        :rtype: str
        """
        chars = list(s)
        vowels = "aeiouAEIOU"
        left = 0
        right = len(s) - 1
        
        while left < right:
            while left < right and chars[left] not in vowels:
                left += 1
            while left < right and chars[right] not in vowels:
                right -= 1
            chars[left], chars[right] = chars[right], chars[left]
            left += 1
            right -= 1
        
        return "".join(chars)
            