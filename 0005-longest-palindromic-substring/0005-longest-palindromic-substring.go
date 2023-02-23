func isPalindrome(s string)bool{
    l,r := 0,len(s)-1
    
    for l <= r{
        if s[l]!=s[r]{
            return false
        }
        l++
        r--
    }
    return true
}

type key struct{
    l,r int
}

func longestPalindrome(s string) string {
    var max string
    solution(s,0,len(s),&max, make(map[key]bool))
    return max
}

func solution(s string, l, r int,max *string, memo map[key]bool){    
    _key := key{l,r}
    if val,e := memo[_key]; e{
        if val && r-l+1 > len(*max)  {
            *max = s[l:r]
        }
        return
    }
    
    if l > r {
        return 
    }
    
    if isPalindrome(s[l:r]) && r-l+1 > len(*max){
        memo[_key] = true
        *max = s[l:r]
        return
    }
    memo[_key] = false
    solution(s,l+1,r,max,memo)
    solution(s,l,r-1,max,memo)
}