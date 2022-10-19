func removeTrash(s string)[]rune{
    res := []rune{}
    var symbol rune
    
    leading := false
    noTrailing :=[]rune{}
    for _,c := range s{
        v := c -'0'        
        if  (c == '+' || c== '-' ) || v >= 0 && v<=9{            
            noTrailing = append(noTrailing, c)
            leading = true
        }else if c == ' ' && !leading{            
            continue
        }else if !leading { // if we don't have a valid char and digit to the left then that means we have an invalid input           
            return []rune{}
        }else{
            break
        }
        
    }
    
    for _,c := range noTrailing{
        v := c - '0'
        
        if (c == '-' || c == '+') {
            if (symbol != 0 || len(res)>0){ // already have a symbol or the symbol comes after a digit
                break
            } else if symbol == 0{                            
                symbol = c
                continue
            }
        }        
        if c ==' '{
            break
        }
        if v >= 0 && v <=9  {
            res = append(res, v)
        }else{
            break
        }
    }
    if symbol == 0{
        symbol= '+'
    }
    return append([]rune{symbol},res...)
}

func myAtoi(s string) int {
    toRune := removeTrash(s)
    if len(toRune) < 1{
        return 0
    }
    var res int
    symbol:= '+'
    if toRune[0] == '+' || toRune[0] == '-'{
        symbol = toRune[0]
        toRune= toRune[1:]
    }
    
    
    for _,c := range toRune {             
        res = res * 10 + int(c)
        fmt.Println(res)
        if symbol == '+' && res > int(1<<31) -1{
        return int(1<<31)-1
        }else if symbol =='-' &&  ^res+1 < int(-1<< 31){
            return int(-1<<31)
        }
    }

    if symbol == '+' && res < 0{
        res = ^res 
    }else if symbol == '-' && res >0 {
        res = ^res+1
    }
    

    if res > int(1<<31) -1{
        return int(1<<31)-1
    }else if res < int(-1<< 31){
        return int(-1<<31)
    }
    
    return res
}