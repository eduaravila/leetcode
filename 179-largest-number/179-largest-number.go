import (
"sort"
"strings"
)
func largestNumber(nums []int) string {    
    numss := make([]string,len(nums))
    for i,num := range nums {
        numss[i] = strconv.Itoa(num)
    }
    sort.Slice(numss,func(a,b int)bool{
        
        return numss[a]+numss[b] > numss[b]+numss[a]                    
    })
    
    
    res := strings.TrimLeft(strings.Join(numss,""),"0")
    
    if len(res) > 0{
        return res
    }
    return "0"
}