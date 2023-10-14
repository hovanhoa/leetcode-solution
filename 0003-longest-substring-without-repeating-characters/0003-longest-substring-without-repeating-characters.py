class Solution(object):
    def lengthOfLongestSubstring(self, s):
            substring = ""
            maxLen = 0
            for i in s:
                if i not in substring:
                    substring = substring + i
                    maxLen = max(maxLen, len(substring))
                else:
                    substring = substring.split(i)[1] + i
            return maxLen
