type Node struct {
    val int
    key int
    next *Node
    prev *Node
}

type List struct {
    head *Node
    tail *Node
}

func newList()*List{
    list := &List{head: &Node{},tail:&Node{} } 
    list.head.next, list.tail.prev = list.tail,list.head
    return list
}

func (l *List) Remove(node *Node){
    prev,next := node.prev,node.next
    
    prev.next = next
    next.prev = prev
    
}



func (l *List) Insert(node *Node){    
    prev := l.tail.prev
    prev.next = node
    l.tail.prev = node
    node.next = l.tail
    node.prev = prev
    
}


type LRUCache struct {
    cache map[int]*Node    
    list *List    
    limit int
}


func Constructor(capacity int) LRUCache {
    return LRUCache{
        cache: make(map[int]*Node),
        list: newList(),
        limit: capacity,
    }
}


func (this *LRUCache) Get(key int) int {
    
    if _,e := this.cache[key]; !e{
        return -1
    }
    this.list.Remove(this.cache[key])
    this.list.Insert(this.cache[key])
    return this.cache[key].val
}


func (this *LRUCache) Put(key int, value int)  {
    if node,e := this.cache[key]; e{
        this.list.Remove(node)        
    }    
    node := &Node{val:value,key:key}
    this.cache[key] = node
    
    this.list.Insert(node)  
    if len(this.cache) > this.limit{        
        nodeToRemove := this.list.head.next
        this.list.Remove(nodeToRemove)
        delete(this.cache,nodeToRemove.key)
    }    
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */