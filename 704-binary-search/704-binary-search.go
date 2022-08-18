func search(nums []int, target int) int {
    if len(nums) < 1 {
        return -1
    }
    l,r := 0,len(nums)-1
    
    for l <= r {
        m := int(uint(l+r)) >>1
        if nums[m] == target {
            return m
        }else if nums[m] < target {
            l = m+1
        }else{
            r = m-1
        }
    }
    return -1
}