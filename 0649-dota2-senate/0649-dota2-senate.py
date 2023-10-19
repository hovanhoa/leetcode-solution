class Solution:
    def predictPartyVictory(self, senate: str) -> str:
        radiantQueue = collections.deque()
        direQueue = collections.deque()

        for i in range(len(senate)):
            if senate[i] == "R":
                radiantQueue.append(i)
            else:
                direQueue.append(i)
        
        while radiantQueue and direQueue:
            radiant, dire = radiantQueue.popleft(), direQueue.popleft()

            if radiant < dire:
                radiantQueue.append(radiant + len(senate))
            else:
                direQueue.append(dire + len(senate))
        
        return "Radiant" if radiantQueue else "Dire"
