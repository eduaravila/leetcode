func getMax(a,b int )int{
    if a > b {
        return a
    }
    return b
}


func deleteAndEarn(nums []int) int {
    set := []int{}
    counter := make(map[int]int)
    
    for _, num := range nums{
        if _,e := counter[num]; !e{
            set = append(set, num)
        }
        counter[num]++
    }
    
    
    sort.Ints(set)
    var a,b int 
    
    for i, num := range set {
        current := num * counter[num]
        
        if i > 0 && set[i-1] == num-1{
            temp := b
            b = getMax(b , current + a)
            a = temp
        }else{
            temp := b
            b += current
            a = temp 
        }
    }
    
    return b
}