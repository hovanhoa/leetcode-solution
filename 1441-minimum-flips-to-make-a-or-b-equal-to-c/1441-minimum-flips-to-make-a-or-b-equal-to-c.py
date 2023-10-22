class Solution:
    def minFlips(self, a: int, b: int, c: int) -> int:
        ans=0
        while a or b or c:
            x,y,z=a%2,b%2,c%2
            print(x,y,z)
            if z==0:ans+=(x+y)
            elif x==0 and y==0:ans+=1
            a,b,c=a//2,b//2,c//2
        return ans