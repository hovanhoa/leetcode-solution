// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
    h := IntHeap(stones)
    heap.Init(&h)
    for h.Len() > 1 {
        x, y := heap.Pop(&h).(int), heap.Pop(&h).(int)
        if x != y {
            heap.Push(&h, x - y)
        }
    }

    if h.Len() == 0 {
        return 0
    }

    return heap.Pop(&h).(int)

}