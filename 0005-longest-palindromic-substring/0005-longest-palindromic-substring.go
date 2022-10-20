
func longestPalindrome(s string) string {
    var res string
    if len(s)  <2{
        return s
    }
    for i := range s{
        solution(s,i,i,&res)
        solution(s,i,i+1,&res)
    }
    return res
}

func solution(s string, l,r int, res *string){
    
    for  l >= 0 && r < len(s)  && s[l] == s[r] {
        
        l--
        r++        
    }
    if r - l > len(*res){
        *res = s[l+1:r]
    }
    
}