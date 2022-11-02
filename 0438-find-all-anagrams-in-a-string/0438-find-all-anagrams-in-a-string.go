func findAnagrams(s string, p string) []int {
    phash := make(map[byte]int)
    res := []int{}
    
    for _,c := range p{
        phash[byte(c)]++
    }
    
    var l,r int 
    for r < len(s) {
        _,e := phash[s[r]]
        if e{
            phash[s[r]]--
        }
        
        for e && phash[s[r]] < 0 && l < r{
            if _,e := phash[s[l]];e{
                phash[s[l]]++
            }
            l++  
        }
        
        if !e {
            for l <= r{             
                if _,e := phash[s[l]];e{
                    phash[s[l]]++
                }
                l++                
            }
            r++
            l=r
            continue            
        }
        
        if r - l +1 == len(p){
            res =append(res,l)
            phash[s[l]]++
            l++
        }      
        
        r++
    }
    
    return res
}