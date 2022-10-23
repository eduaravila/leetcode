type node struct{
    prev *node
    next *node
    val int
    key int
}

type nodeList struct{
    head *node
    tail *node
}

type LRUCache struct {
    values map[int] *node
    list *nodeList
    capacity int
}

func (n *nodeList) add(nnode *node){
    nnode.prev = nil
    nnode.next = nil
    prev,next := n.head,n.head.next
    
    nnode.prev= prev
    nnode.next= next
    prev.next = nnode
    next.prev = nnode
}

func (n *nodeList) remove(nnode *node){
    prev,next := nnode.prev,nnode.next
    prev.next = next 
    next.prev = prev
}

func Constructor(capacity int) LRUCache {
    head,tail := &node{}, &node{}
    head.next = tail
    tail.prev = head
    list := &nodeList{head:head,tail:tail}
    return LRUCache{values:make(map[int]*node),list: list,capacity:capacity}
}


func (this *LRUCache) Get(key int) int {
    n,e:= this.values[key]
    if !e {
        return -1
    }
    
    this.list.remove(n)    
    this.list.add(n)    
    return n.val    
}


func (this *LRUCache) Put(key int, value int)  {
    
    n,e := this.values[key]
    
    if e {
        this.list.remove(n)
    }

    if this.capacity < 1 && !e{
        temp := this.list.tail.prev
        this.list.remove(temp)
        delete(this.values,temp.key)
    }
    if this.capacity > 0 && !e{         
        this.capacity--
    }
    
    nnode := &node{val: value,key: key}    
    
    this.list.add(nnode)
    this.values[key] = nnode
    
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

// ["LRUCache","put","put","put","put","get","get"]
//[[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]
// 2:3
// 4:1