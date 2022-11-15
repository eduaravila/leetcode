func removeKdigits(num string, k int) string {
    stack := []rune{}
    for _, c:= range num{
        for len(stack) > 0 && c < stack[len(stack)-1] && k > 0{
            stack = stack[:len(stack)-1]
            k--
        }
    
        stack = append(stack, c)
    }
    if k > 0{
        stack = stack[:len(stack)-k]
    }
    res := strings.TrimLeft(string(stack),"0")
    if len(res) < 1{
        return "0"
    }
    return res
}

//1432219
//1219
//3,2,1,0

//10200
//0200 
//1