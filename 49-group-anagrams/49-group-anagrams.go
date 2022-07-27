import (
"strings"
"sort"
)
func groupAnagrams(strs []string) [][]string {
    groups := make(map[string][]string)
    for _,str := range strs{
        s := []rune(str)
        sort.Slice(s, func(a,b int) bool {return s[a] < s[b]})
        key := string(s)
        groups[key] = append(groups[key],str)
    }
    res := make([][]string, 0)
    for _,group := range groups{
        res = append(res, group)
    }
    return res
}