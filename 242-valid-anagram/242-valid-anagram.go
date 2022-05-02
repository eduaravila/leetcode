func isAnagram(s string, t string) bool {
    // count symbols on s and store them on a hash table
    // iterate through t and check if the current symbol exist on the count hash table
    // if true,
    // check if you can reduce 
    // reduce 1 from the hash table
    // if the counter is already 0 then that means t is not an anagram
    // check if the counter if empty meaning we have the same symbols on both strings
    counter := make(map[rune]int)
    for _,symbol := range s{
        counter[symbol]++
    }
    
    for _,symbol := range t {
        if _,e := counter[symbol]; e {            
            counter[symbol]--
            if counter[symbol] < 1 {
                delete(counter,symbol)
            }
        }else{
            return false
        }
    }
    
    return len(counter) == 0
}