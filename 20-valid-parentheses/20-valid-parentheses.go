
func isCompatible(opening rune, closing rune)bool{
    switch(opening){
        case '(':
        if closing == ')'{
            return true
        }        
        case '{':
        if closing == '}'{
            return true
        }
        case '[':
        if closing == ']'{
            return true
        }        
    }
    return false
}

func isOpening(symbol rune)bool{
    return symbol == '(' || symbol == '[' || symbol == '{'
}

func isValid(s string) bool {
    // opening char must have ending char
    // compare compatible symbols
    // add to stack opening symbols
    // every time we find a closing symbol, compare with peak from stack
    
    opening_stack := make([]rune,0)
    
    for _, symbol := range s{
        if isOpening(symbol){
            opening_stack = append(opening_stack,symbol)
        }else if len(opening_stack) > 0 && isCompatible(opening_stack[len(opening_stack)-1], symbol){
            opening_stack = opening_stack[:len(opening_stack)-1]
        }else{
            return false;
        }
        
    }
    
    return len(opening_stack) < 1
}