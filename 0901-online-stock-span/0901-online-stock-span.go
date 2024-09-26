type StockSpanner struct {
    stack []int
}


func Constructor() StockSpanner {
    return StockSpanner{[]int{}}
}


func (this *StockSpanner) Next(price int) int {
    this.stack = append(this.stack, price)
    ans := 0
    for i := 0; i < len(this.stack); i++ {
        if this.stack[i] > price {
            ans = 0
        } else {
            ans++
        }
    }

    return ans
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */