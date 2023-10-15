class Solution(object):
    def largestAltitude(self, gain):
        """
        :type gain: List[int]
        :rtype: int
        """
        maxPoint = 0
        for indx in range(len(gain)):
            maxPoint = max(maxPoint, sum(gain[:indx+1]))

        return maxPoint
