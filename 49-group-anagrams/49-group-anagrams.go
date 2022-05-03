func getHashKey(str string)string{
    counter := make([]rune, 26)
    for _,symbol := range str{
        counter[symbol -'a']++
    }
    return string(counter)
}

func groupAnagrams(strs []string) [][]string {
    // create a hashkey with the words, meaning same symbols in apparence and amount 
    // results in the same hash key 
    // use an array of the size of the english alphabet 
    // the groups hash table will save an array of strings
    // add the values to the final response
    
    groups := make(map[string][]string)
    for _, str := range strs {
        key := getHashKey(str)
        groups[key]= append(groups[key],str)
    }
    
    res := [][]string{}
    for _, group := range groups {
        res = append(res,group)
    }
    
    return res 
}