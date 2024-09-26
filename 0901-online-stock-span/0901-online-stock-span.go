type StockSpanner struct {
    stack []int
}


func Constructor() StockSpanner {
    return StockSpanner{[]int{}}
}


func (this *StockSpanner) Next(price int) int {
    ans := 1
    for i := len(this.stack) - 1; i >= 0; i-- {
        if this.stack[i] > price {
            break
        }

        ans++
    }

    this.stack = append(this.stack, price)

    return ans
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */