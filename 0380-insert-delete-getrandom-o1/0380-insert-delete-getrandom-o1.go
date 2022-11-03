type RandomizedSet map[int]bool


func Constructor() RandomizedSet {
    return make(map[int]bool)
}


func (this *RandomizedSet) Insert(val int) bool {
    if _,e := (*this)[val];e{
        return false
    }
    (*this)[val]= true
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if _,e := (*this)[val];!e{
        return false
    }
    delete(*this,val)
    return true
}


func (this *RandomizedSet) GetRandom() int {
    values := []int{}
    for val:= range *this{
        values = append(values,val)
    }
    return values[rand.Intn(len(values))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */