func minWindow(s string, t string) string {
    if len(t) > len(s){
        return *new(string)
    }
   
    count := make(map[rune]int)
    for _,val := range t{
        count[val]++
        
    }
    
    l := 0
    intf :=int(^uint(0)>>1)
    res := []int{l,intf}
    subCount := make(map[rune]int)
    
    matches:=0
    for r := range s{ // O(n)
        k := rune(s[r])    
        subCount[k]++
        if val, e := count[k]; e && subCount[k] == val {                        
            matches++            
        }
        
        for matches == len(count){     
            if r - l  < res[1] - res[0] {
                res = []int{l,r}    
            }
            k = rune(s[l])
            subCount[k]--
            if val, e := count[k]; e && subCount[k] < val{
                matches--
            }
            l++
        }
              
    }
    if res[1] == intf{
        return *new(string)
    }
    return s[res[0]:res[1]+1]
}

// r keeps going if the substring is not complete
// l moves to shrink the result
// count the chars in t 
// if the current subSting has the required chars 
// enter inner loop to try to reduce the size if the substring and response
// if the deleted element was necesary in the substring, exit the loop 
// response need l and r
