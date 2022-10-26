func longestCommonPrefix(strs []string) string {
    res := []string{}
    
    word := strs[0]
    for i, c := range word {
        for _,str := range strs{
            
            if  i > len(str)-1 || rune(str[i]) != c {
                return strings.Join(res,"")
            }
        }
        res = append(res,string(c))
    }
    return strings.Join(res,"")
}