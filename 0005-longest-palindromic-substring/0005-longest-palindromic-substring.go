func longestPalindrome(s string) string {
    var res string
    
    for i := range s{
        getBiggestPalindrome(s,i,i, &res)
        getBiggestPalindrome(s,i,i+1, &res)
    }
    return res
}

func getBiggestPalindrome(strn string, s,e int,  res *string){
    var pal string
    for s > -1 && e < len(strn) && strn[s] == strn[e]{
        pal = strn[s:e+1]
        s--
        e++
        
    }
    
    if len(*res) < len(pal) {
        *res = pal
    }    
}