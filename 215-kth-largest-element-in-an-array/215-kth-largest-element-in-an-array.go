import ("container/heap")

type IntHeap []int


func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


func findKthLargest(nums []int, k int) int {
    
    minH := &IntHeap{}

    heap.Init(minH)

    for _,num := range nums{
        heap.Push(minH,num)
    }

    for minH.Len() > k {
        heap.Pop(minH)
    }
    
    return heap.Pop(minH).(int)
}