import ("container/heap")

type minH []int

func (m minH) Less(a,b int)bool { return m[a] < m[b] }
func (m minH) Len()int { return len(m) }
func (m minH) Swap(a,b int) { m[a],m[b] = m[b],m[a] }

func (m *minH) Pop()interface{}{
    n := len(*m)-1
    temp := (*m)[n]
    *m = (*m)[:n]
    return temp
}

func (m *minH) Push(val interface{}){
    *m = append(*m, val.(int))
}

type KthLargest struct {
    size int
    minH *minH
}

func Constructor(k int, nums []int) KthLargest {
    
    minH := &minH{}
    *minH = append(*minH, nums...)
    heap.Init(minH)
    
    for  minH.Len() > k{                        
        heap.Pop(minH)        
    }
    return KthLargest{
        size: k,
        minH: minH,
    }
}


func (this *KthLargest) Add(val int) int {
    heap.Push(this.minH,val)
    if this.minH.Len() > this.size{
        heap.Pop(this.minH)
    }
    return (*this.minH)[0]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */