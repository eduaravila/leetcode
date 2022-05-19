import 
(
    "fmt"
)
func longestConsecutive(nums []int) int {
    counter := make(map[int]bool)
    
    for _,num := range nums {
        counter[num]=true
    }
    
    uf := make(map[int]int)
    
    for _, num := range nums{
        if _,e := counter[num+1];e{
            uf[num]= num+1        
        }else{
            uf[num] = num
        }        
    }
    
    for num := range uf {
        group := []int{}
        current := num
        parent := uf[current]
        for current != parent {
            group= append(group,current)
            current = uf[current]
            parent = uf[current]
        }
        group = append(group,current)
        
        for _,elem := range group{
            uf[elem] = current    
        }
        
        group=[]int{}
    }   
    
    groups := make(map[int]int)
    for _,group := range uf{
        groups[group]++
    }
    mx := 0
    
    for _,val := range groups{
        if val > mx {
            mx= val
        }
    }
    return mx
}