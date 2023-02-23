
type key struct {
    l,r int
}
func longestPalindrome(s string) string {
    var max string
    for i := range s{
        getPalindrome(s,i, i, &max)
        getPalindrome(s,i,i+1, &max)
    }
    return max
}


func getPalindrome(s string, l,r int, max *string){        
    var templ, tempr int
    for l >= 0 && r <len(s) && s[l]==s[r] {
        templ, tempr = l, r                                           
        l--
        r++
    }
    
    if tempr - templ +1 > len(*max){        
        *max = s[templ:tempr+1]
    }
}
