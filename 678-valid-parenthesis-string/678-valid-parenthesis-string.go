func checkValidString(s string) bool {
    var min,max int
    
    for _,c := range s{
        if c =='('{
            min++
            max++
        }else if c == '*'{
            max++
            min--
        }else{
            min--
            max--
        }
        if max < 0{
            return false
        }
        if min < 0{
            min =0
        }    
    }
    
    return min == 0
}