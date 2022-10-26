func groupAnagrams(strs []string) [][]string {
    mapping := make(map[string][]string)
    
    for _, str := range strs{
        counter := make([]int,'z'-'a'+1)
        for _,c := range str{
            counter[c-'a']++
        }
        key := fmt.Sprintf("%v", counter)
        mapping[key] = append(mapping[key],str)
    }
    res := [][]string{}
    
    for _,group := range mapping{
        res =append(res,group)
    }
    return res
}