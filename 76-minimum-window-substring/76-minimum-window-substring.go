func minWindow(s string, t string) string {
    if len(t) > len(s){
        return *new(string)
    }
   
    count := make(map[rune]int)
    for _,val := range t{
        count[val]++
        
    }
    
    l,r := 0,0
    res := []int{l,int(^uint(0)>>1)}
    subCount := make(map[rune]int)
    
    matches:=0
    for r < len(s){
        k := rune(s[r])        
        if val, e := count[k]; e  {            
            subCount[k]++
            if subCount[k] == val{
                matches++
            }
        }
        
        for matches == len(count){     
           
            if r - l  < res[1] - res[0] {
                res = []int{l,r+1}    
            }
            k := rune(s[l])
            subCount[k]--
            if val, e := count[k]; e && subCount[k] < val{
                matches--
            }
            l++
        }
        r++        
    }
    if res[1] == int(^uint(0)>>1){
        return *new(string)
    }
    return s[res[0]:res[1]]
}

// r keeps going if the substring is not complete
// l moves to shrink the result
// count characters of t using a array from 0 to 26