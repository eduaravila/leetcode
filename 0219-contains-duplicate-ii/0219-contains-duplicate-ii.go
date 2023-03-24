func containsNearbyDuplicate(nums []int, k int) bool {
    l,r := 0,0
    counter := make(map[int]int)
    
    for  r < len(nums) {
        num := nums[r]
        counter[num]++
        
        if counter[num] > 1 {
            return true
        }
        
        if (r-l+1) > k {
            counter[nums[l]]--
            l++            
        }
        
        r++
    }
    
    return false
}
