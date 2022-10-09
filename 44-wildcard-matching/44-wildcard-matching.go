func isMatch(s string, p string) bool {    
   var is,ip,ast,match int
    ast = -1
    for is < len(s){
        if ip < len(p) && (p[ip] == '?' || p[ip] ==s[is]){
            is++
            ip++
            continue
        }
        if ip < len(p) && p[ip] == '*'{
            ast = ip
            match = is
            ip++
            continue
        }
        
        if ast > -1 {
            ip = ast+1
            match++
            is = match
            continue
        }
        
        return false
    }
    fmt.Println(ip,is,ast)
    for _,e := range p[ip:]{
        if e !='*'{
            return false
        }
    }
    return true
}

