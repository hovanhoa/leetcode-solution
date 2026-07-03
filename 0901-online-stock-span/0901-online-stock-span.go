type StockSpanner struct {
    stack []int
}


func Constructor() StockSpanner {
    return StockSpanner{
        stack: []int{},
    }
}


func (this *StockSpanner) Next(price int) int {
    this.stack = append(this.stack, price)
    n := 0
    for i := len(this.stack)-1; i >=0; i-- {
        if this.stack[i] <= price {
            n += 1
        } else {
            break
        }
    }

    return n
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */