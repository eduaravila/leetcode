import ("container/heap")

type maxHeapElement struct {
    val float64
    index int
}

type maxHeap []maxHeapElement

func (m maxHeap) Less(a,b int)bool {return m[a].val < m[b].val}
func (m maxHeap) Len()int { return len(m) }
func (m maxHeap) Swap(a,b int) { m[a], m[b]  = m[b], m[a] }

func (m *maxHeap) Push(val interface{}) {     
    *m = append(*m, val.(maxHeapElement))
}

func (m *maxHeap) Pop() interface{}{
    temp := (*m)[m.Len()-1]    
    *m = (*m)[:m.Len()-1]
    return temp
}


func kClosest(points [][]int, k int) [][]int {
    distances := &maxHeap{}
    heap.Init(distances)
    
    for i, value := range points{
        result := maxHeapElement{val: math.Sqrt( math.Pow( float64( 0 - value[0]  ), 2 )  + math.Pow( float64( 0 - value[1] ), 2 ) ), index:i}
        heap.Push(distances,result)
        
    }
    res := [][]int{}
    
    for i := 0 ; i < k ; i++{
        res = append(res, points[heap.Pop(distances).(maxHeapElement).index] )
    }

    return res
}