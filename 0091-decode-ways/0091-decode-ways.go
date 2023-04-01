func numDecodings(s string) int {
    return solution(s,0,make(map[int]int))
}

func solution(s string, l int,memo map[int]int)int{
        
    if l > len(s)-1 {
         return 1 // last value we assume we found a valid way
    }
    if val,e := memo[l]; e{
        return val
    }
   
    i, _ := strconv.Atoi(s[l:l+1])
    if i == 0{
        return 0
    }
    
    var res int 
    if i == 1 && l < len(s)-1 {
        res += solution(s,l+2, memo)
    }
    
    
    if l < len(s)-1 {
        nextI, _ := strconv.Atoi(s[l+1:l+2]) 
        if i == 2 && nextI <= 6{
            res += solution(s,l+2, memo)
        }
    }
   
    res += solution(s, l+1, memo)
    memo[l] = res
    return res
}

/*
AAJF


    11106
    
 var l, r -> save the current index for s 
 
 if r > len(s) -1 {
     return 1 // last value we assume we found a valid way
 }
 
 check if current is in range 
 
 if s[l:r] < 'a' || s[l:r]  > 'z' {
 // 0
 // 35
    return 0    // this is invalid
 }
 
return solution(s, r,r+1)  + solution(s, r, r+2)

"226"


2 , 26,
2, 2, 6 


22, 6




202

20,2




*/