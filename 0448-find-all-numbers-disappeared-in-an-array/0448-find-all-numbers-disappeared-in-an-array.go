func findDisappearedNumbers(nums []int) []int {
    counter := make(map[int]bool)
    
    for _, num := range nums{
        counter[num] = true
    }
    res := []int{}
    current := 1
    for current <= len(nums){        
        if _,e:= counter[current];!e{
            res =append(res,current)    
        }
        current++
    }
    return res
}