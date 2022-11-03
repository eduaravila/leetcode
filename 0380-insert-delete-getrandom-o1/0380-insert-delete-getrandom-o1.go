type RandomizedSet struct{
    set map[int]int
    values []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        set: make(map[int]int),
        values: []int{},
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _,e := this.set[val];e{
        return false
    }
    n := len(this.values)
    this.set[val]= n
    this.values = append(this.values,val)
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if _,e := this.set[val];!e{
        return false
    }
    n := this.set[val]    
    temp := this.values[len(this.values)-1]
    this.values[n], this.values[len(this.values)-1] = this.values[len(this.values)-1], this.values[n]    
    this.values = this.values[:len(this.values)-1]    
    this.set[temp] = n
    delete(this.set,val)    
    return true
}


func (this *RandomizedSet) GetRandom() int {        
    return this.values[rand.Intn(len(this.values))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */