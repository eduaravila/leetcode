// This example demonstrates an integer heap built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type element struct {
    val byte
    times int
}
type IntHeap []element

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i].times > h[j].times }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(element))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func leastInterval(tasks []byte, n int) int {
    h := &IntHeap{}
    heap.Init(h)
    times := make(map[byte]int)
    for _, c := range tasks{
        times[c] +=1
    }
    
    for val, t := range times{
        heap.Push(h, element{val:val,times: t})
    }
    
    var res int
    for h.Len() > 0{
        

        popped:= []element{}
        current := heap.Pop(h).(element)
        popped = append(popped,current)
        var i int
        for i < n && h.Len() > 0{
            i++            
            current = heap.Pop(h).(element)
            popped = append(popped,current)         
        }
        
        if h.Len() < 1 && popped[0].times == 1{
            res += len(popped)
            break
        }
        for _,p := range popped{
            if p.times < 2{
                continue
            }
            heap.Push(h,element{val:p.val,times:p.times-1})
        }
        res += n+1

    }
    return res
}
    

// 
//A: 6 | 6 * n + 6 = 18
//B: 1
//C: 1
//E: 1
//F: 1
//G: 1



// create heap with struct { c char, times int}
// add c, remove from heap, continue with the next value of the heap, update or re-insert the values to the heap
// a,b,c
// a,d,e
// a,f,g
// a,
// a,
// a 

// total: 10