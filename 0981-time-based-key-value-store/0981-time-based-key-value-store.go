type data struct{
    value string
    timestamp int
}

type TimeMap struct {
    keys map[string]int
    values [][]data
}


func Constructor() TimeMap {
    return TimeMap{
        keys: make(map[string]int),
        values: [][]data{},
    }    
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    if _,e := this.keys[key];!e{
        this.keys[key] = len(this.values)
        this.values = append(this.values, []data{})
    }         
    index := this.keys[key]
    this.values[index] = append(this.values[index],data{value,timestamp})
}


func (this *TimeMap) Get(key string, timestamp int) string {
    if _, e:=this.keys[key];!e{
        return ""
    }   
    res := -1
    index := this.keys[key]
    l ,r := 0,len(this.values[index])-1
    for l <= r {
        m := (l+r) >> 1
        if this.values[index][m].timestamp <= timestamp{
            res = m
            l= m+1
        }else{
            r=m-1
        }
    }
    
    if res < 0 {
        return ""
    }
    
    return this.values[index][res].value
    
    
}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

//["TimeMap","set","set","get","get","get","get","get"]
//[[],["love","high",10],["love","low",20],["love",5],["love",10],["love",15],["love",20],["love",25]]
//[null,null,null,"low","high","low","low","low"]
//[null,null,null,"","high","high","low","low"]