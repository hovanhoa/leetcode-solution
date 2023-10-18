class Solution:
    def predictPartyVictory(self, senate: str) -> str:
        n = len(senate)
        radiant, dire = [], []
        for i, v in enumerate(senate):
            if v == "R":
                radiant.append(i)
            else:
                dire.append(i)
        
        while radiant and dire:
            r, d  = radiant.pop(0), dire.pop(0)
            if r < d:
                radiant.append(r + n)
            else:
                dire.append(d + n)

        return "Radiant" if radiant else "Dire"
