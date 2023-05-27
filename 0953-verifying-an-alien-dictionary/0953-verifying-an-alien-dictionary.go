func isAlienSorted(words []string, order string) bool {
    dict_indexes := make(map[rune]int)
    
    for i, c := range order{
        dict_indexes[c]=i+1
    }
    dict_indexes[' '] = 0
    
    var i int
    for i < len(words)-1{

        a, b := words[i], words[i+1]
        
        var aindex,bindex int
        for aindex < len(a) || bindex < len(b){
            var aletter, bletter rune
            if aindex < len(a){
                aletter = rune(a[aindex])
                aindex++
                
            }
            
            if bindex < len(b){
                bletter = rune(b[bindex])
                bindex++
            }           
            
            if dict_indexes[aletter] < dict_indexes[bletter]{
                break
            }
            if dict_indexes[aletter] > dict_indexes[bletter]{
                return false
            }
            
            
        }
        i++
    }
    
    return true
}

/*

   compare two words 
    
    for l <lenb(a) || r < len(b)
        
        if dict[a[l]] > dict[b[r]]        
            return false
    
*/