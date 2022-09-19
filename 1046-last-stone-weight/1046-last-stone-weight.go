
import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
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


func lastStoneWeight(stones []int) int {
    minH := &IntHeap{}
    *minH = append([]int{},stones...)
    heap.Init(minH)
    for minH.Len() > 1 {
        max,min := heap.Pop(minH).(int), heap.Pop(minH).(int)
        result := max - min 
        if max > min {
            heap.Push(minH, result)
        }
        
    }
    
    heap.Push(minH,0)
    return heap.Pop(minH).(int)
}