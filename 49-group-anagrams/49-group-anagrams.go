import (
    "sort"
    "bytes"
)

func groupAnagrams(strs []string) [][]string {
    if len(strs) < 1{
        return [][]string{}
    }
    
    groups := make(map[string][]string)
    res := [][]string{}
    for _, str := range strs{
        ordered := bytes.NewBufferString(str)
        sort.Slice(ordered.Bytes(), func(a,b int)bool{
            return ordered.Bytes()[a] < ordered.Bytes()[b]
        })
        
        groups[ordered.String()] = append(groups[ordered.String()],str)
    }
    
    for _,group := range groups{
        res = append(res,group)
    }
    
    return res
}