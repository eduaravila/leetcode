func containsDuplicate(nums []int) bool {
    set := make(map[int]bool)
    
    for _,val := range nums {
        if _,e := set[val]; e{
            return true
        }
        set[val] = true       
    }
    
    return false
}