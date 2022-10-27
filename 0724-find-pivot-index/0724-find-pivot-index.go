func pivotIndex(nums []int) int {
    left_right := make([]int,len(nums))
    right_left := make([]int,len(nums))
    l,r := 0, len(nums)-1
    var s1,s2 int
    for l < len(nums){
        left_right[l] = s1
        right_left[r] = s2
        s1 += nums[l]
        s2 += nums[r]        
        r--
        l++
    }
    
    l = 0
    for l < len(nums) {
        if left_right[l] == right_left[l]{
            return l
        }
        
        l++
    }
    return -1
}