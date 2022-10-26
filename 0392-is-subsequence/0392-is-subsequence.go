func isSubsequence(s string, t string) bool {
    a,b := 0,0
    for b < len(t){
        if a > len(s)-1{
            return true
        }
        if s[a] == t[b]{
            a++
        }
        b++
    }
    return a > len(s)-1
}