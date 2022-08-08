func isValidClose(r rune, stack []rune)bool{
    if len(stack) < 1{
        return false
    }
    last := stack[len(stack)-1]    
    if (r == '}' && last == '{') || (r == ')' && last == '(') || (r == ']' && last == '['){
        return true
    }
    return false
}

func isValid(s string) bool {
    cStack := []rune{}
    for _, r := range s{
        if r == '(' || r == '[' || r == '{' {
            cStack = append(cStack,r)
        }else if isValidClose(r,cStack){
            cStack = cStack[:len(cStack)-1]
        }else{
            return false
        }
    }
    return len(cStack) == 0
}