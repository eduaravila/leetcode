import ("container/heap")

type intHeap []int

func (i intHeap) Less(a,b int)bool {return i[a] < i[b]}
func (i intHeap) Len()int {return len(i)}
func (i intHeap) Swap(a,b int) { i[a], i[b] = i[b], i[a] }

func (i *intHeap) Pop() interface{} {
    temp := (*i)[i.Len()-1]
    *i = (*i)[:i.Len()-1]
    return temp
}

func (i *intHeap) Push(val interface{}){
    *i = append(*i, val.(int))
}

type KthLargest struct {
    intH *intHeap
    size int
}


func Constructor(k int, nums []int) KthLargest {
    minH := &intHeap{}
    *minH = append([]int{}, nums...)
    heap.Init(minH)
    for minH.Len() > k{
        heap.Pop(minH)
    }
    return KthLargest{
        intH: minH,
        size: k,
    }
}


func (this *KthLargest) Add(val int) int {
    heap.Push(this.intH, val)
    if this.intH.Len() > this.size{
        heap.Pop(this.intH)
    }
    
    
    return (*this.intH)[0]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */