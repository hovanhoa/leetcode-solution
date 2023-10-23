class SmallestInfiniteSet:
    def __init__(self):
        self.c = 1
        self.s = set()
    def popSmallest(self):
        if self.s:
            r = min(self.s)
            self.s.remove(r)
            return r
        else:
            self.c += 1
            return self.c - 1
    def addBack(self, num):
        if self.c > num:
            self.s.add(num) 