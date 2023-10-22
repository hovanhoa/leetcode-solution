class StockSpanner:

    def __init__(self):
        self.memo = []
        self.timestamp = 0
    
        
    def next(self, price: int) -> int:
        self.timestamp += 1
        record = {'ts':self.timestamp, 'price':price}

        while(self.memo and price >= self.memo[-1]['price']):
            self.memo.pop()
        
        self.memo.append(record)
        if len(self.memo) > 1:
            return (self.memo[-1]['ts'] - self.memo[-2]['ts'])
        
        else:
            return self.timestamp
        


# Your StockSpanner object will be instantiated and called as such:
# obj = StockSpanner()
# param_1 = obj.next(price)