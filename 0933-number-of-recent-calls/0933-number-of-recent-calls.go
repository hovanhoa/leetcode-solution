type RecentCounter struct {
    q []int
}


func Constructor() RecentCounter {
    return RecentCounter{
        q: []int{},
    }
}


func (this *RecentCounter) Ping(t int) int {
    cnt := 0
    for _, v := range this.q {
        if v >= t - 3000 && v <= t {
            cnt += 1
        }
    }

    this.q = append(this.q, t)

    return cnt + 1
}


/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */