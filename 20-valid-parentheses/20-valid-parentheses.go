

func isOpening(symbol rune)bool{
    return symbol == '(' || symbol == '[' || symbol == '{'
}

func antonymOf(symbol rune)rune{
    switch (symbol){
        case '(':
            return ')'
        case '[':
            return ']'
        case '{':
            return '}'
    }
   return 0
}

func isValid(s string) bool {
    // opening char must have ending char
    // compare compatible symbols
    // add to stack opening symbols
    // every time we find a closing symbol, compare with peak from stack
    
    opening_stack := make([]rune,0)
    
    for _, symbol := range s{
        n := len(opening_stack)-1
        if isOpening(symbol){            
            opening_stack = append(opening_stack, antonymOf(symbol))
        }else if n >= 0 && opening_stack[n]  == symbol {
            opening_stack = opening_stack[:n]
        }else{
            return false;
        }
        
    }
    
    return len(opening_stack) < 1
}