class Solution:
    def reverseVowels(self, s: str) -> str:
        vowels = "aeiouAEIOU"
        l, r = 0, len(s) - 1
        nums = list(s)

        while l < r:
            while l < r and nums[l] not in vowels:
                l += 1
            while l < r and nums[r] not in vowels:
                r -= 1
            if l < r:
                nums[l], nums[r] = nums[r], nums[l]

            l += 1
            r -= 1

        
        return "".join(nums)
