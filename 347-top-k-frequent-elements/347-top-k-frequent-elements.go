import (
    "container/heap"
    "fmt"
)


type element struct {
    num int
    freq int
}

type FreqHeap []*element

func (F FreqHeap) Len()int {return len(F)}
func (F FreqHeap) Less(a,b int)bool {return F[a].freq < F[b].freq }
func (F FreqHeap) Swap(a,b int) {    
    F[a], F[b] = F[b], F[a]
}

func (F *FreqHeap) Push(elem interface{}) {
    *F = append(*F, elem.(*element))
}

func (F *FreqHeap) Pop() interface{}{
    
    old := *F // value
    n := len(old)
    temp := old[n-1]
    *F = old[:n-1]
    
    return temp
}

func topKFrequent(nums []int, k int) []int {
    counter := make(map[int]int)
    
    for _, num := range nums {
        counter[num]++
    }
    minHeap := &FreqHeap{}
    heap.Init(minHeap)
    
    for num, freq := range counter {
        heap.Push(minHeap, &element{num:num, freq:freq} )
        
        if minHeap.Len() > k {
            
            heap.Pop(minHeap)
        }
    }
    res := []int{}
    for i := 0 ; i < k ; i++{
        res = append(res, heap.Pop(minHeap).(*element).num)
    }
    
    return res
}