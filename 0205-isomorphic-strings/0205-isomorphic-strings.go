func isIsomorphic(s string, t string) bool {
    s1,s2 := make(map[byte]byte), make(map[byte]byte)
    var i int
    for i < len(s){     
        v1,e1 := s1[t[i]] 
        v2,e2 := s2[s[i]] 
        if (e1 || e2) && ( v1 != s[i] ||  v2 != t[i] ) {
            return false
        }else{
            s1[t[i]] = s[i]
            s2[s[i]] = t[i]
        }
        i++
    }
    return true
}

// e:a, g: d
// a:e, d: g
