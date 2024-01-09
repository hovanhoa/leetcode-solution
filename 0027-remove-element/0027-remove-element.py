class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        indexCount = 0
        for i, value in enumerate(nums):
            if val != value:
                nums[indexCount], nums[i] = nums[i], nums[indexCount]
                indexCount+=1

        return indexCount 
        