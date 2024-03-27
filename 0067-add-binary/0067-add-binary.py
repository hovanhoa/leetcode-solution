class Solution(object):
    def addBinary(self, a, b):
        """
        :type a: str
        :type b: str
        :rtype: str
        """
        i, j = len(a) - 1, len(b) - 1
        carry = 0
        res = ""
        while i >= 0 or j >= 0 or carry:
            num = carry
            if i >= 0:
                num += int(a[i])
                i -= 1
            if j >= 0:
                num += int(b[j])
                j -= 1
            res += str(num % 2)
            carry = num // 2
        return "".join(reversed(res))