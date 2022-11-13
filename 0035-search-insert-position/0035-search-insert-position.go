func searchInsert(nums []int, target int) int {
    l,r := 0,len(nums)-1
    
    for l <= r{
        m := (int(l+r) >>1)
        
        if nums[m] == target {
            return m
        }
        if nums[m] < target{
            l = m+1
        }else{
            r = m-1
        }        
    }
    return l
}


// 1,3,5,6  -> 7
// l 0 2 3 4
// r 3 3 3 3
// m 1 2 3 
