type MinStack [][]int


func Constructor() MinStack {
    return [][]int{}    
}

func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}

func (this *MinStack) Push(val int)  {
    inf := int(^uint(0)>>1)
    cMin := inf
    if len(*this) > 0 {
        cMin = this.GetMin()
    }
    
    *this = append(*this, []int{val, getMin(cMin,val)})
}


func (this *MinStack) Pop()  {
    *this = (*this)[:len(*this) -1]
}


func (this *MinStack) Top() int {
    return (*this)[len(*this)-1][0]
}


func (this *MinStack) GetMin() int {
    return (*this)[len(*this)-1][1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */