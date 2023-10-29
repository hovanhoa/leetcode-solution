class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        indexMap = {}
        for i in range(len(nums)):
            if (target - nums[i]) not in indexMap:
                indexMap[nums[i]] = i
            else:
                return [indexMap[target - nums[i]], i]
        
        return [-1, -1]
