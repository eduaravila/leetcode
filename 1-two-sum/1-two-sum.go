func twoSum(nums []int, target int) []int {
    counter := make(map[int][]int)
    
    for idx, num := range nums {
        counter[num] = append(counter[num],idx)
    }
    
    
    res := []int{}
    for num, idxs := range counter { 
        remaining := target - num 
        res = append(res, idxs[0])
        if _, e :=counter[remaining];e {
            if remaining == num && len(idxs) > 1 {
                res = append(res,idxs[1])
                return res
            }else if remaining != num {
                res = append(res,counter[remaining][0])   
                return res
            }            
        }
        res = []int{}
    }
    
    return res
}