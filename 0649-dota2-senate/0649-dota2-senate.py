class Solution:
    def predictPartyVictory(self, senate: str) -> str:
        radiant = collections.deque()
        dire = collections.deque()

        for i in range(len(senate)):
            if senate[i] == "R":
                radiant.append(i)
            else:
                dire.append(i)
        
        while radiant and dire:
            first, second = radiant.popleft(), dire.popleft()

            if first < second:
                radiant.append(first + len(senate))
            else:
                dire.append(second + len(senate))

        return "Radiant" if radiant else "Dire"