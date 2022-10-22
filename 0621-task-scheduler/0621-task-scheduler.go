// This example demonstrates an integer heap built using the heap interface.
package main

import (
	"container/heap"
	"fmt"
)


type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j]}
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

func leastInterval(tasks []byte, n int) int {
    h := &IntHeap{}
    heap.Init(h)
    times := make(map[byte]int)
    for _, c := range tasks{
        times[c] +=1
    }
    
    for _, t := range times{
        heap.Push(h, t)
    }
    
    var res int
    var current int
    for h.Len() > 0{
        if (*h)[0] == 1{
            return res + h.Len()
        }
        popped:= []int{}
        current = heap.Pop(h).(int)
        popped = append(popped,current)
        var i int
        for i < n && h.Len() > 0{
            i++            
            current = heap.Pop(h).(int)
            popped = append(popped,current)         
        }
        for _,p := range popped{
            if p ==1 {
                continue
            }
            heap.Push(h,p-1)
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