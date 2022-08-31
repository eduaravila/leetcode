func permute(nums []int) [][]int {
    if len(nums) == 1{
        return [][]int{{nums[0]}}
    }
    res := [][]int{}
    for range nums {
        num := nums[0]        
        nums = append(nums[:0],nums[1:]...)
        perms := permute(nums)
        for i := range perms{  
            perms[i] = append(perms[i],num)
        }
        nums =append(nums,num)
        res = append(res, perms...)
    }
    return res
}

