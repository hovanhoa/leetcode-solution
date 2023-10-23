import heapq
class Solution:
    def maxScore(self, nums1: List[int], nums2: List[int], k: int) -> int:
        lst=list(zip(nums2,nums1))
        lst.sort(key=lambda x:(-x[0],-x[1]))
        flst=[]
        heapq.heapify(flst)
        i=0
        sm=0
        ef=float("infinity")
        prd=float("-infinity")
        while i<k:
            x=lst.pop(0)
            heapq.heappush(flst,x[1])
            ef=min(ef,x[0])
            sm+=x[1]
            i+=1
        prd=max(prd,sm*ef)
        while lst:
            x=heapq.heappop(flst)
            sm-=x
            y=lst.pop(0)
            heapq.heappush(flst,y[1])
            ef=min(ef,y[0])
            sm+=y[1]
            prd=max(prd,sm*ef)
        return prd