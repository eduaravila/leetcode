func getMax(a,b int)int{
    if a > b {
        return a
    }
    return b
}
func longestConsecutive(nums []int) int {
    counter := make(map[int]bool)
    
    for _,num := range nums {
        counter[num]=true
    }
    
    var res int
    for _,num := range nums {
        var groupl int = 0
        if _,e := counter[num-1]; !e{
            _, e := counter[num+groupl]
            for e {                
                groupl++
                _, e = counter[num+groupl]
            }
        }
        res = getMax(groupl,res)        
    }
    return res
}