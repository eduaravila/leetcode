
/*
if it starts at 0 it returns 0 
  0 -> 0
  01 -> 0
  
only 1 and 2 accept two digits 


only 1 accept digits from 0 to 9

2 from 0 to 6


if current has two digits return 1

12 
i

12
 i
 
after all the other conditiosn have passed 
 
if i > len(s){
return 1
}

*/

func numDecodings(s string) int {

    
    return solution(s, 0, make(map[int]int))
}

func solution(s string, l int, memo map[int]int) int{
    
    if val,e := memo[l];e{
        return val
    }
    
    if l > len(s)-1{        
        return 1
    }
 
    lnum, _ :=strconv.Atoi(string(s[l]))
    
    if lnum == 0 {
        return 0
    }
    
    var res int 
    if lnum == 1 && l < len(s)-1 {
        res += solution(s,l+2,memo)        
    }
    
    if l < len(s)-1 {
        rnum, _ := strconv.Atoi(string(s[l+1:l+2]))
        if lnum == 2 &&  rnum <=6 {        
            res +=solution(s,l+2, memo)
        }            
    }
    
    res += solution(s,l+1, memo)
    memo[l] = res
    return res   
}