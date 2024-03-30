class Solution(object):
    def findDifference(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: List[List[int]]
        """
        l1 = [x for x in nums1 if x not in nums2]
        l2 = [x for x in nums2 if x not in nums1]
        return [list(set(l1)), list(set(l2))]
