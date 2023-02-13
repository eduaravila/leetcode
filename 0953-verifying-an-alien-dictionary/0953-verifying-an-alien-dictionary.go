func isAlienSorted(words []string, order string) bool {
    dict := make(map[rune]int)
    
    for i,c := range order{
        dict[c] = i
    }
    
    
    for x:= 0; x <len(words)-1;x++{
        w1,w2 := words[x],words[x+1]
        for j := 0 ; j < len(w1) ; j++{
            // second word is smaller
            if j ==len(w2){
                return false
            }
            
            if w1[j] != w2[j]{
                if dict[rune(w1[j])] > dict[rune(w2[j])]{
                    return false
                }
                break
            }
            
        }
    }
    return true
}