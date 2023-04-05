func lengthOfLIS(nums []int) int {
    sub := []int{nums[0]}
    
    
    for _, num := range nums {
        if num <= sub[len(sub)-1] {
            l,r := 0,len(sub)-1
            for l <= r {
                m := int((l+r) >> 1)
                if sub[m] >= num {
                    r = m-1
                }else{
                    l = m+1
                }
            }
            sub[l] = num            
        }else{
            sub = append(sub,num)
        }
    }
    
    return len(sub)
}