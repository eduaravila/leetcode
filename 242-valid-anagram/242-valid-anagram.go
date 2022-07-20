func isAnagram(s string, t string) bool {
    counterA := make(map[rune]int)
    counterB := make(map[rune]int)
    
    for _,val := range s{
        counterA[val]++
    }
    for _,val := range t{
        if _,e := counterA[val]; !e{
            return false
        }
        counterB[val]++
    }
    
    for letter,counter := range counterA{
        if counterB[letter] != counter{
            return false
        }
    }
    
    return true
}