class Solution(object):
    def gcdOfStrings(self, str1, str2):
        """
        :type str1: str
        :type str2: str
        :rtype: str
        """
        if str1 == str2:
            return str1

        if len(str1) < len(str2):
            return self.gcdOfStrings(str2, str1)
        
        if str1.startswith(str2):
            return self.gcdOfStrings(str2, str1[len(str2):])

        return ""
