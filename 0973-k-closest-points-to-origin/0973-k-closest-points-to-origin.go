import (
    "container/heap"
    "math"
)

type distancesHeap [][]int

func (h distancesHeap) Less(a,b int)bool {
    distance1 := math.Sqrt(math.Pow(0-float64(h[a][0]), 2) + math.Pow(0-float64(h[a][1]), 2))
    distance2 := math.Sqrt(math.Pow(0-float64(h[b][0]), 2) + math.Pow(0-float64(h[b][1]), 2))
    return distance1 < distance2
}
func (h distancesHeap) Len()int{return len(h)}
func (h distancesHeap) Swap(a,b int){h[a],h[b] = h[b],h[a]}
func (h *distancesHeap) Push(val interface{}){
    *h = append(*h, val.([]int))
}
func (h *distancesHeap) Pop() interface{}{
    temp := (*h)[h.Len()-1]
    *h = (*h)[:h.Len()-1]
    return temp
}

func kClosest(points [][]int, k int) [][]int {
    distances := &distancesHeap{}
    heap.Init(distances)
    
    for _,point := range points {
        heap.Push(distances,point)
    }
    
    res := [][]int{}
    for k > 0 {
        res = append(res,heap.Pop(distances).([]int))
        k--
    }
    return res
}