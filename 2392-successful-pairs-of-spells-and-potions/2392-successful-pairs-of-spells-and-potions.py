class Solution:
    def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:

        n, ans = len(potions), []
        potions.sort()
        
        for s in spells:
            num = success//s + bool(success%s)
            idx = bisect_left(potions, num)
            ans.append(n-idx)

        return ans