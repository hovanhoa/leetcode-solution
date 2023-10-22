class Solution:
    def findMinArrowShots(self, points: List[List[int]]) -> int:
        #sort on the basis of ending time
        points = sorted(points, key=lambda x: x[1])
        currPointer = points[0][1]
        arrows = 1

        for i in range(len(points)):
            if points[i][0] > currPointer: #if start of i > last pointer pointing at
                currPointer = points[i][1]
                arrows += 1
        return arrows