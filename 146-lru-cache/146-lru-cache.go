type Node struct {
    Val int
    Key int
    Next *Node
    Prev *Node
}

type List struct {
    Head *Node
    Tail *Node    
}

func NewList()*List{    
    head, tail := &Node{}, &Node{}
    head.Next = tail
    tail.Prev = head
    list := &List{Head:head, Tail:tail}
    return list
}

func (l *List) Remove(node *Node){    
    prev,next := node.Prev, node.Next
    prev.Next = next
    next.Prev = prev
    node.Next = nil
    node.Prev = nil
}

func (l *List) Insert(node *Node){
    prev := l.Tail.Prev
    node.Prev = prev
    node.Next = l.Tail
    prev.Next = node
    l.Tail.Prev = node
    
}

type LRUCache struct {
    values map[int]*Node
    list *List
    size int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{
        size: capacity,
        values: make(map[int]*Node),
        list: NewList(),
    }
}

func (this *LRUCache) Get(key int) int {
    node, e := this.values[key]
    if !e {
        return -1
    }
    this.list.Remove(node)
    this.list.Insert(node)    
    return node.Val
}

func (this *LRUCache) Put(key int, value int) {
    if node,e := this.values[key]; e{
        this.list.Remove(node)
    }
    
    newNode := &Node{Val:value, Key:key}
    this.list.Insert(newNode)
    this.values[key] = newNode
    
    if len(this.values) > this.size{
        leastUsedNode := this.list.Head.Next
        this.list.Remove(leastUsedNode)
        delete(this.values,leastUsedNode.Key)        
    }
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */