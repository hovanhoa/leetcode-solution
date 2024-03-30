class Solution(object):
    def uniqueOccurrences(self, arr):
        """
        :type arr: List[int]
        :rtype: bool
        """
        res = []
        for i in list(set(arr)):
            res.append(arr.count(i))

        return len(res) == len(set(res))