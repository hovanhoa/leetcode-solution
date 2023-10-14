class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        index_map = dict()
        for index, value in enumerate(nums):
            if target - value in index_map:
                return [index_map[target - value], index]
            else:
                index_map[value] = index
