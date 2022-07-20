func isAnagram(s string, t string) bool {
    
    if len(s) != len(t){
        return false
    }
    
    counterA := make(map[byte]int)
    counterB := make(map[byte]int)
    
    for i := range s{
        counterA[s[i]]++
        counterB[t[i]]++
    }
    
    for letter,counter := range counterA{
        if counterB[letter] != counter{
            return false
        }
    }
    
    return true
}