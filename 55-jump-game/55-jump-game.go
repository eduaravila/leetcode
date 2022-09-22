func getMax(a,b int)int{
    if a > b {
        return a
    }
    return b
}

func canJump(nums []int) bool {
    furtest := 0 
    
    for i,jumps := range nums{
        if furtest < i {
            return false
        }
        furtest = getMax(i+jumps,furtest)
    }
    
    return true
}