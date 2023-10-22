import collections
import itertools


class UnionFind(object):
    def __init__(self):
        self.set = {}
        self.rank = collections.Counter()

    def find_set(self, x):
        xp, xr = self.set.setdefault(x, (x, 1.0))
        if x != xp:
            pp, pr = self.find_set(xp)  # path compression.
            self.set[x] = (pp, xr*pr)  # x/pp = xr*pr
        return self.set[x]

    def union_set(self, x, y, r):
        (xp, xr), (yp, yr) =  map(self.find_set, (x, y))
        if xp == yp:
            return False
        if self.rank[xp] < self.rank[yp]:  # union by rank
            # to make x/yp = r*yr and merge xp into yp
            # => since x/xp = xr, we can merge with xp/yp = r*yr/xr 
            self.set[xp] = (yp, r*yr/xr)
        elif self.rank[xp] > self.rank[yp]:
            # to make y/xp = 1/r*xr and merge xp into yp
            # => since y/yp = yr, we can merge with yp/xp = 1/r*xr/yr 
            self.set[yp] = (xp, 1.0/r*xr/yr)
        else:
            # to make y/xp = 1/r*xr and merge xp into yp
            # => since y/yp = yr, we can merge with yp/xp = 1/r*xr/yr 
            self.set[yp] = (xp, 1.0/r*xr/yr)
            self.rank[xp] += 1 
        return True

    def query_set(self, x, y):
        if x not in self.set or y not in self.set:
            return -1.0
        (xp, xr), (yp, yr) = map(self.find_set, (x, y))
        return xr/yr if xp == yp else -1.0


class UnionFindPathCompressionOnly(object):
    def __init__(self):
        self.set = {}

    def find_set(self, x):
        xp, xr = self.set.setdefault(x, (x, 1.0))
        if x != xp:
            pp, pr = self.find_set(xp)  # path compression.
            self.set[x] = (pp, xr*pr)  # x/pp = xr*pr
        return self.set[x]

    def union_set(self, x, y, r):
        (xp, xr), (yp, yr) =  map(self.find_set, (x, y))
        if xp == yp:
            return False
        # to make x/yp = r*yr and merge xp into yp
        # => since x/xp = xr, we can merge with xp/yp = r*yr/xr 
        self.set[xp] = (yp, r*yr/xr)
        return True

    def query_set(self, x, y):
        if x not in self.set or y not in self.set:
            return -1.0
        (xp, xr), (yp, yr) = map(self.find_set, (x, y))
        return xr/yr if xp == yp else -1.0


class Solution(object):
    def calcEquation(self, equations, values, queries):
        """
        :type equations: List[List[str]]
        :type values: List[float]
        :type queries: List[List[str]]
        :rtype: List[float]
        """
        union_find = UnionFind()
        for (a, b), k in itertools.izip(equations, values):
            union_find.union_set(a, b, k)
        return [union_find.query_set(a, b) for a, b in queries]
