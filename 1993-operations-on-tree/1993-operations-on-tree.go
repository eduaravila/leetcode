type stateKey struct {
    state bool
    user int
}

func (s *stateKey) isLocked()bool{
    return s.state 
}

func (s *stateKey) isValidUser(userid int)bool{
    return s.user == 0 || userid == s.user
}

func (s *stateKey) lock(userid int){
    s.state = true
    s.user = userid
}

func (s *stateKey) unlock(){
    s.state = false
    s.user = 0
}



type LockingTree struct {
    parent []int
    state []*stateKey
    parents map[int][]int
}


func Constructor(parent []int) LockingTree {
    state := []*stateKey{}    
    parents := make(map[int][]int)

        
    for i,node :=  range parent{
        state = append(state, &stateKey{})
        
        if _,e := parents[node];!e{
            parents[node] = []int{}
        }
        parents[node] = append(parents[node],i)
    }
    return LockingTree{
        parent:parent,
        state: state,
        parents:parents,
    }
}



func (this *LockingTree) Lock(num int, user int) bool {
    if this.state[num].isLocked() || !this.state[num].isValidUser(user){
        return false
    }
    
    this.state[num].lock(user)
    return true
}


func (this *LockingTree) Unlock(num int, user int) bool {
    if !this.state[num].isLocked() || !this.state[num].isValidUser(user){
        return false
    }
    this.state[num].unlock()
    return true
}


func (this *LockingTree) Upgrade(num int, user int) bool {
    if this.state[num].isLocked() || this.hasLockedAncestors(num) || !this.hasLockedDescendants(num){        
        return false        
    }
    this.state[num].lock(user)
    for _, index := range this.getDecendants(num) {
        this.state[index].unlock()
    }
    return true
}

func (this *LockingTree) getDecendants(num int) []int {    
    children := []int{}    
    queue := []int{num}
    for len(queue) > 0 {
        n := len(queue)        
        for _,element := range queue {
            children = append(children,this.parents[element]...)
            queue= append(queue,this.parents[element]...)
        }
        queue = queue[n:]
    }
    
    return children
}


func (this *LockingTree) hasLockedDescendants(num int) bool{
    
    for _, i := range this.getDecendants(num){
        if this.state[i].isLocked(){
            return true
        }
    }
    return false
}

func (this *LockingTree) hasLockedAncestors(num int) bool{
    
    current := this.parent[num]
    for current != -1{        
        if this.state[current].isLocked(){
            return true
        }
        current = this.parent[current]
    }
    
    return false
}

/**
 * Your LockingTree object will be instantiated and called as such:
 * obj := Constructor(parent);
 * param_1 := obj.Lock(num,user);
 * param_2 := obj.Unlock(num,user);
 * param_3 := obj.Upgrade(num,user);
 */