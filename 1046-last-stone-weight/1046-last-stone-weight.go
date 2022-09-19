
import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
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


func lastStoneWeight(stones []int) int {
    minH := &IntHeap{}
    
    heap.Init(minH)
    
    for len(stones) > 1 {
        for _, stone := range stones{
            heap.Push(minH,stone)            
            stones = stones[1:]
            if minH.Len() > 2{
                stones =append(stones, heap.Pop(minH).(int))
            }
        }
        min, max := heap.Pop(minH), heap.Pop(minH)
        result := max.(int) - min.(int)
        if result < 1{
            continue
        }
        stones = append(stones,result)
    }
    if len(stones ) < 1{
        return 0
    }
    return stones[0]
}