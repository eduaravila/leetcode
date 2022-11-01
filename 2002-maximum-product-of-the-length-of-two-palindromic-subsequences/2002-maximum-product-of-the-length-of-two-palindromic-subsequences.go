func getMax(a,b int)int{
    if a > b {
        return a
    }    
    return b
}

func isPalindrome(s string)bool{
    if len(s)<1{
        return false
    }
    l,r := 0, len(s)-1
    for l < r {
        if s[l] != s[r]{
            return false
        }
        l++
        r--
    }
    return true
}

func maxProduct(s string) int {    
    sequences := make(map[int]int)
    for i:=1 ; i < (1 << len(s)) ;i++{
        sets := []rune{}
        for j := 0 ; j < len(s) ; j++{
            if i&(1 << j) > 0 {
                sets= append(sets, rune(s[j]))
            }
        }
        if isPalindrome(string(sets)){            
            sequences[i] = len(sets)
        }
    }
    var max int
    for x,val1 := range sequences{
        for y,val2 := range sequences{
            if x&y == 0{                            
                max = getMax(max,val1 * val2)
            }
        }
    }
    return max
}