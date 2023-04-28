func longestNiceSubstring(s string) string {
    l,r := solution(s, 0, len(s))
    return s[l:r]
}


func solution(s string, l, r int)(int,int){
    substring := make(map[string]bool)
    
    for i := l ; i < r; i++ {
        substring[string(s[i])]= true
    }
    
    for i := l ; i < r; i++ {
        c := string(s[i])
        
        _, upperExist := substring[strings.ToUpper(c)]
        _,lowerExist := substring[strings.ToLower(c)]
        
        if upperExist && lowerExist {
            continue
        }
        
        a1,b1 := solution(s, i+1,r)
        a2,b2 := solution(s, l, i)
        
        if b1 - a1 > b2 - a2 {
            return a1,b1
        }
        return a2,b2
    }
    return l,r
}