func getMin(vals ...int)int{
    res := vals[0]
    for _, c := range vals{
        if c < res{
            res = c
        }
    }
    return res
}

func flip(a string) string{
    if a =="1"{
        return "0"
    }
    return "1"
}

func minFlips(s string) int {
    p1,p2 := "1","0"
    alt1,alt2 := []string{},[]string{}
    var d1,d2,l,r int
    res := int(^uint(0)>>1)
    sarr := strings.Split(s,"")
    sarr = append(sarr,sarr...)
    
    for r < len(s) *2{
        c := sarr[r]
        alt1 =append(alt1,p1)
        alt2 =append(alt2,p2)
        if c != p1{
            d1++
        }
        if c != p2{
            d2++
        }
        
        if (r - l + 1) > len(s){
            if alt1[l] != sarr[l]{
                d1--
            }
            if alt2[l] != sarr[l]{
                d2--
            }
            l++
        }
        if (r-l+1) == len(s){
            res = getMin(d1,d2,res)                        
        }
        p1 = flip(p1)
        p2 = flip(p2)
        r++
    }
    
   
    
    return res
}

// 111000111000
// 010101010101
// 101010101010

// d1 4
// d2 2


