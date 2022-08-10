func characterReplacement(s string, k int) int {
    l,r, max:=0,0,0
    count := make(map[byte]int)
    res :=0
    for r < len(s){
        key := s[r]
        count[key]++        
        if count[key]  > max{
            max =count[key]
        }
        
        if ((r - l) +1 ) - max > k {
            count[s[l]]--            
            l++
        }
        if ((r - l) +1 ) > res{            
            res = r-l+1
        }
        r++
    }
    return res
}