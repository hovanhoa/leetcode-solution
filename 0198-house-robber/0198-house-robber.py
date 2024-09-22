class Solution:
    def rob(self, nums: List[int]) -> int:
        if len(nums) <= 2:
            return max(nums)
        memorize = [0]*len(nums)
        memorize[-1] = nums[-1]
        memorize[-2] = nums[-2]
        for i in range(len(nums)-3, -1, -1):
            memorize[i] = max(nums[i], nums[i] + max(memorize[i+2:]))
        return max(memorize[0], memorize[1])