type FreqStack struct {
    values map[int][]int
    set []int
    counter map[int]int
}


func Constructor() FreqStack {
    return FreqStack{
        values: make(map[int][]int),
        set: []int{},
        counter: make(map[int]int),
    }
}


func (this *FreqStack) Push(val int)  {
    this.counter[val]++
    if len(this.set) < 1 || this.counter[val] > this.set[len(this.set)-1]{
        this.set = append(this.set, this.counter[val])
    }
    this.values[this.counter[val]] = append(this.values[this.counter[val]], val)
}


func (this *FreqStack) Pop() int {
    max := this.set[len(this.set)-1]
    val := this.values[max][len(this.values[max])-1]
    this.values[max] = this.values[max][:len(this.values[max])-1] 
    this.counter[val]--
    if len(this.values[max]) < 1{
        delete(this.values,max)
        this.set = this.set[:len(this.set)-1]
    }
    return val
}


/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */

// {1:[1,2]}
// 1
// counter {1:3,2:1}
//["FreqStack","push","push","push","push","pop", "pop", "push", "push", "push", "pop", "pop", "pop"]
//[[],[1], [1], [1], [2], [], [], [2], [2], [1], [], [], []]