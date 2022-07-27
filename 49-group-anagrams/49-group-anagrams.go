import (
"strings"
"sort"
)
func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)
    for _,str := range strs{
        sorted := strings.Split(str,"")        
        sort.Strings(sorted)
        key := fmt.Sprintf("%v", sorted)
        
        groups[key] = append(groups[key],str)
    }
    res := make([][]string, 0)
    for _,group := range groups{
        res = append(res, group)
    }
    return res
}