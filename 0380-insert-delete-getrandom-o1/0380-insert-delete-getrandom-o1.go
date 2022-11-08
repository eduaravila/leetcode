type RandomizedSet struct {
    values map[int]int
    indexes []int
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        values: make(map[int]int),
        indexes: []int{},
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if _,e := this.values[val]; e{
        return false
    }
    this.values[val] = len(this.indexes)
    this.indexes = append(this.indexes,val)
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if _, e:= this.values[val];!e{
        return false
    }
    index := this.values[val]
    n := len(this.indexes)-1
    this.indexes[index],this.indexes[n] = this.indexes[n], this.indexes[index]    
    this.values[this.indexes[index]] = index
    
    delete(this.values,val)
    this.indexes = this.indexes[:n]
    
    return true
}


func (this *RandomizedSet) GetRandom() int {
    return this.indexes[rand.Intn(len(this.indexes))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */