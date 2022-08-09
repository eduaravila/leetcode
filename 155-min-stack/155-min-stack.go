type element struct {
    val int
    min int
}
type MinStack []element



func Constructor() MinStack {
    return []element{}
}


func (this *MinStack) Push(val int)  {
    if len(*this) < 1 || val < this.GetMin(){
        *this = append(*this, element{val: val, min:val})    
    }else{
        *this = append(*this, element{val:val, min: this.GetMin()})
    }
}


func (this *MinStack) Pop()  {
    *this = (*this)[:len(*this)-1]
}


func (this *MinStack) Top() int {
    return (*this)[len(*this)-1].val
}


func (this *MinStack) GetMin() int {
    if len(*this) < 1{
        return 0
    }
    return (*this)[len(*this)-1].min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */