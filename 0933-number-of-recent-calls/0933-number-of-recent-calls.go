type RecentCounter struct {
    ping []int
}


func Constructor() RecentCounter {
    return RecentCounter{}
}


func (this *RecentCounter) Ping(t int) int {
    ans := 1
    for i := len(this.ping) - 1; i >= 0; i-- {
        if this.ping[i] >= t - 3000 {
            ans++
        } else {
            break
        }
    }

    this.ping = append(this.ping, t)
    return ans
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */