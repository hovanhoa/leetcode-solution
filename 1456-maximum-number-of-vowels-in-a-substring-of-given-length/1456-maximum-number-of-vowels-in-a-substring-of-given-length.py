class Solution(object):
    def maxVowels(self, s, k):
        """
        :type s: str
        :type k: int
        :rtype: int
        """
        vowels = "aeiou"
        cnt = ans = sum(s[i] in vowels for i in range(k))

        for i in range(k, len(s)):
            cnt += (s[i] in vowels) - (s[i - k] in vowels)
            ans = max(ans, cnt)

        return ans
                