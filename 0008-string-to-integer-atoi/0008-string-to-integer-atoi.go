func getMin(a,b int) int{
    if a < b{
        return a
    }
    return b
}

func myAtoi(s string) int {
    
    s = strings.TrimLeftFunc(s, func(r rune) bool {
		return unicode.IsSpace(r)
	})
    if len(s) <1 {
        return 0
    }
    l := strings.IndexFunc(s,func(c rune)bool{
        return unicode.IsNumber(c)   
    })
    
    
    symbol := '+'
    if s[0] == '-' || s[0] == '+' {
        symbol = rune(s[0])
        s = s[1:]            
    }
    
    letterI := strings.IndexFunc(s, func(c rune)bool{
        return unicode.IsLetter(c)
    })
    spaceI := strings.IndexFunc(s, func(c rune)bool{
        return unicode.IsSpace(c)
    })
    if l == -1 || (letterI != -1 && letterI < l) || (spaceI != -1 && spaceI < l) {
        return 0
    }
    end_index := len(s)
    
    if letterI != -1{        
        end_index = getMin(end_index,letterI)
    }
    if spaceI != -1{        
        end_index = getMin(end_index,spaceI)
    }


    var res int
    low_limit,high_limit := int(-1<<31), int(1<<31)-1
    for _,c := range s[:end_index]{
        v := int(c - '0')
        if v < 0 || v > 9{
            break
        }
        res = (res * 10) + v
        if symbol == '-' && ^res + 1 < low_limit{
            return low_limit
        }else if symbol == '+' && res > high_limit{
            return high_limit
        }
    }
    
    if symbol == '-'{
        res = ^res +1
    }
    return res   
}