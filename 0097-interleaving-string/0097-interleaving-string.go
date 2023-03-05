type key struct {
    a,b int
}

func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1) + len(s2) != len(s3){
        return false
    }
    dp := [][]bool{}
    for i := 0 ; i <= len(s1); i++ {
        dp = append(dp, make([]bool,len(s2)+1))        
    }
    
    dp[len(s1)][len(s2)] = true
    
    for l := len(s1); l >= 0 ; l--{
        for r := len(s2) ; r >= 0 ; r--{            
            if l < len(s1) && s1[l] == s3[l+r] && dp[l+1][r]{
                dp[l][r] = true
            }
            if r < len(s2) && s2[r] == s3[l+r] && dp[l][r+1]{
                dp[l][r] = true
            }
        }
                                        
    }
    
    return dp[0][0]
}


/*
  a a b c c 
d  
b
b
c
a
            
    a a d b b c b c a c
    
    
    lastCommon = 
    s1 2
    s2 1

    
    
    a a d 
    
*/

/*
func solution(s1 string, s2 string, s3 string, i,a,b int, memo map[key]bool) bool{
    k := key{a,b}
    if val,e := memo[k];e{
        return val
    }
    
    if s1 == "" {
        return true
    }    

    if a > len(s2) -1 && b > len(s3) -1 && i > len(s1) -1 {
        return true
    }
    
    if i >len(s1) -1{
        return false
    }
    
    if a < len(s2) && b < len(s3) && s2[a] == s1[i] && s3[b]== s1[i] {
        res :=  solution(s1,s2,s3,i+1,a+1,b, memo) || solution(s1,s2,s3,i+1,a,b+1,memo)
        memo[k] = res
        return res
    }
    
    if a < len(s2) && s2[a] == s1[i]{
        return solution(s1,s2,s3,i+1,a+1,b, memo)
    }
    
    if b < len(s3) && s3[b] == s1[i]{
        return solution(s1,s2,s3,i+1,a,b+1, memo)
    }
    
    memo[k] = false
    
    return false
}
*/