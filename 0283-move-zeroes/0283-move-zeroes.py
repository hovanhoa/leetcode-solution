class Solution(object):
    def moveZeroes(self, nums):
        """
        :type nums: List[int]
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        indx = 0           
        for i in range(len(nums)):
            if nums[i] != 0:
                nums[indx], nums[i] = nums[i], nums[indx]
                indx += 1
        
        return
