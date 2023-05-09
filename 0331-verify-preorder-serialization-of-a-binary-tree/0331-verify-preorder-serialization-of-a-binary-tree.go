func isValidSerialization(_preorder string) bool {
    stack := []string{}
    preorder := strings.Split(_preorder,",") 
    for _, c := range preorder {
        stack= append(stack, c)
        
        for len(stack) > 1 && stack[len(stack)-1] == "#" && stack[len(stack)-2] == "#"{
            stack = stack[:len(stack)-2]
            if len(stack)< 1{
                return false
            }
            stack[len(stack)-1] = "#"
        }
    }
    if len(stack) == 1 && stack[0]== "#"{
        return true
    }
    return false
}