class Solution:
    def asteroidCollision(self, asteroids: List[int]) -> List[int]:
        stack = []
        
        for i in asteroids:
            if i > 0:
                stack.append(i)
            else:
                while stack and stack[-1] > 0 and stack[-1] < -i:
                    stack.pop()
                if not stack or stack[-1] < 0:
                    stack.append(i)
                elif stack and stack[-1] == -i:
                    stack.pop()
        
        return stack
