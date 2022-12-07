type Node struct{
    Prev *Node
    Next *Node
    Val int
}

type MyCircularQueue struct {
    head *Node 
    tail *Node
    capacity int
    size int
}


func Constructor(k int) MyCircularQueue {
    head  := &Node{Val: -1}
    tail := &Node{Val: -1}
    head.Prev,head.Next  = tail, tail
    tail.Prev,tail.Next = head, head
    return MyCircularQueue{
        head:head,
        tail:tail,
        capacity:k,
        size:0,
    }
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull(){
        return false
    }
    prev,next := this.tail.Prev ,this.tail
    node := &Node{Val:value}    
    prev.Next,next.Prev = node,node
    node.Prev,node.Next = prev,next
    this.size++
    return true
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty(){
        return false
    }
    node:= this.head.Next
    prev,next  :=node.Prev,node.Next
    prev.Next, next.Prev = next, prev    
    this.size--
    return true
}


func (this *MyCircularQueue) Front() int {
    return this.head.Next.Val
}

func (this *MyCircularQueue) Rear() int {
    return this.tail.Prev.Val
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.size <1
}

func (this *MyCircularQueue) IsFull() bool {
    return this.size >= this.capacity
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

