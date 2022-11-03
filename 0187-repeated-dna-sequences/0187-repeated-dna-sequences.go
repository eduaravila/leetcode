func findRepeatedDnaSequences(s string) []string {    
    sequence := []rune{}
    sequences := make(map[string]int)
    if len(s) < 10{
        return []string{}
    }
    for i := 0; i < 10 ; i++{
        sequence = append(sequence,rune(s[i]))
    }
    
    sequences[string(sequence)]++
    
    r := 10
    for r < len(s){        
        sequence = sequence[1:]
        sequence = append(sequence,rune(s[r]))
        r++
        sequences[string(sequence)]++
    }
    
    res := []string{}
    for val,times := range sequences{
        if times > 1{
            res = append(res,val)            
        }
    }
    return res
}