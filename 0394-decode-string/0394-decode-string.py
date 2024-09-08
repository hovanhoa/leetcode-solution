class Solution(object):
    def decodeString(self, s):
        """
        :type s: str
        :rtype: str
        """
        stack = []
        curStr = ""
        curNum = 0

        for c in s:
            if c.isdigit():
                curNum = curNum * 10 + int(c)
            elif c == "[":
                stack.append((curStr, curNum))
                curStr = ""
                curNum = 0
            elif c == "]":
                preStr, num = stack.pop()
                curStr = preStr + curStr * num
            else:
                curStr += c
            
        return curStr
        