import (
"strings"
)
func isPalindrome(input string) bool {
    s := strings.ToLower(input)
    
    l,r := 0, len(s)-1
    for l < r {
        if (s[r] <'a' || s[r] >'z') && (s[r] <'0' || s[r] >'9') {
            r--
            continue
        }
        if (s[l] <'a' || s[l] >'z') && (s[l] <'0' || s[l] >'9') {
            l++
            continue
        }
        
        
        if s[l] != s[r]{
            return false
        }
        
        r--
        l++
    }
    
    return true
}