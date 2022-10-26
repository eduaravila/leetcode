import (
"strings"
"sort"
)
func groupAnagrams(strs []string) [][]string {
    mapping := make(map[string][]string)
    
    for _, str := range strs{
        counter := make([]int,'z'-'a'+1)
        for _,c := range str{
            counter[c-'a']++
        }
        as := strings.Split(str,"")
        sort.Strings(as)
        key := strings.Join(as,"")
        mapping[key] = append(mapping[key],str)
    }
    res := [][]string{}
    
    for _,group := range mapping{
        res =append(res,group)
    }
    return res
}