func minMovesToMakePalindrome(s string) int {
    odd := 0
    sarr := strings.Split(s,"")
    counter := make(map[string]int)
    
    for _,r :=range sarr{
        counter[r]++
    }
    for _,v := range counter {
        if v % 2 !=0{
            odd++
        }
    }
    
    if odd > 1 {
        return 0
    }
    
    lm,rm := 0, len(s)-1
    res := 0
    for lm < rm{
        if sarr[lm] == sarr[rm]{
            rm--
            lm++
            continue
        }
        r:= rm
        for sarr[lm] != sarr[r] && r > lm{
            r--
        }
        if lm == r {
            sarr[lm],sarr[lm+1] = sarr[lm+1],sarr[lm]
            res++
        }else{
            for r < rm{
                sarr[r],sarr[r+1] = sarr[r+1],sarr[r]
                res++
                r++
            }
            lm++
            rm--
        
        }
        
        
    }
    return res
}