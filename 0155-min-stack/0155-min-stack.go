type MinStack [][]int

func getMin(a,b int)int{
    if a < b{
        return a
    }
    return b
}
func Constructor() MinStack {
    return [][]int{}
}

func (this *MinStack) len()int{
    return len(*this)
}

func (this *MinStack) Push(val int)  {
    currentMin := int(^uint(0)>>1)
    if this.len() > 0{
        currentMin = this.GetMin()
    }    
    *this = append(*this,[]int{val,getMin(currentMin,val)})
}


func (this *MinStack) Pop()  {
    *this = (*this)[:this.len()-1]
}


func (this *MinStack) Top() int {
    return (*this)[this.len()-1][0]
}


func (this *MinStack) GetMin() int {
    return (*this)[this.len()-1][1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */