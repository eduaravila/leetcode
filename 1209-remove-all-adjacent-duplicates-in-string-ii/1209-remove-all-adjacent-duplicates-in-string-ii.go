
func removeDuplicates(s string, k int) string {    
    stack := []rune{}
    count := []int{}
    for _, c:= range s{
        if len(stack) > 0 && stack[len(stack)-1] == c{
            count[len(count)-1]++
        }else{            
            count = append(count,1)
        }
        stack = append(stack,c)
        for len(stack) > 0 && count[len(count)-1] == k {            
            stack = stack[:len(stack)-k]
            count = count[:len(count)-1]
        }
        
    }
    return string(stack)
}


