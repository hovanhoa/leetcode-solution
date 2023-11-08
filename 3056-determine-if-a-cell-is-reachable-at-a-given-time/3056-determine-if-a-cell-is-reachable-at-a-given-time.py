class Solution:
    def isReachableAtTime(self, sx, sy, fx, fy, t):
        # If the starting and destination cells are the same, check if the time limit allows for it.
        if sx == fx and sy == fy:
            return t > 1 or t == 0

        # Calculate the height and width differences between the starting and destination cells.
        height_difference = abs(sy - fy)
        width_difference = abs(sx - fx)

        # Calculate the extra time needed to cover the extra moves.
        extra_time = abs(height_difference - width_difference)

        # Check if the total time (min of height and width differences + extra time) is within the time limit.
        return (min(height_difference, width_difference) + extra_time) <= t